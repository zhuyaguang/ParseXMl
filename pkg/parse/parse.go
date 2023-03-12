package parse

import (
	"encoding/json"
	"fmt"
	"github.com/beevik/etree"
	"github.com/colinmarc/hdfs"
	"io/ioutil"
	"log"
	"patentExtr/pkg"
	"patentExtr/pkg/Hadoop"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var NUM = 0
var MAXFILE = 500000
var ERRORNUN = 0

// Par0Xml 主要解析 30-S 类型的专利
func Par0Xml(xmlPath, output string, patentIndex int, client *hdfs.Client) error {
	// log.Println("xml path ,output path type:", xmlPath, output, patentIndex)

	// 得到 XML 文件的名称，比如：CN302021000671538CN00003070960400SDBPZH20220201CN00M
	fileName := filepath.Base(xmlPath)
	fileName = strings.Split(fileName, ".")[0]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlPath); err != nil {
		log.Println(err, "解析失败手动处理========", xmlPath)
		return nil
	}
	doc.SelectElement("PatentDocumentAndRelated")
	//log.Println("ROOT element:", root.Tag)

	patentOBJ := pkg.Patent{}

	resNameE := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/DesignTitle")
	if resNameE != nil {
		patentOBJ.Name = resNameE.Text()
		patentOBJ.ApplicationNO = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PatentNumber").Text()
		patentOBJ.ApplicationDate = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicationReference/DocumentID/Date").Text()

		publicationCN := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/WIPOST3Code").Text()
		publicationNO := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/DocNumber").Text()
		publicationKind := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Kind").Text()
		patentOBJ.PublicationNO = publicationCN + publicationNO + publicationKind
		patentOBJ.PublicationDate = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Date").Text()

		applicationAddressE := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/CorrespondenceAddress/AddressBook/Address/Text")
		if applicationAddressE == nil {
			applicationAddressE = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicantDetails/Applicant/AddressBook/Address/Text")
		}
		patentOBJ.ApplicantAddress = applicationAddressE.Text()

		patentOBJ.Applicant = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicantDetails/Applicant/AddressBook/Name").Text()

		// 发明人
		applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/DesignBibliographicData/DesignerDetails/Designer")
		applicationInventors := ""
		for _, v := range applicationInventorsList {
			g := v.FindElement("./AddressBook/Name").Text()
			applicationInventors = applicationInventors + g + ","
		}
		patentOBJ.Inventors = applicationInventors

		// 摘要
		abstractList := doc.FindElements("./PatentDocumentAndRelated/DesignBriefExplanation/Paragraphs")
		abstract := ""
		for _, v := range abstractList {
			g := v.Text()
			abstract = abstract + g
		}
		patentOBJ.Abstract = abstract

		patentOBJ.MainClassificationNO = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ClassificationLocarno/MainClassification").Text()

		AgencyE := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/Agency/AddressBook/OrganizationName")
		if AgencyE != nil {
			patentOBJ.Agency = AgencyE.Text()
		}

		AgentE := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/AddressBook/Name")
		if AgentE != nil {
			patentOBJ.Agent = AgentE.Text()
		}

		// NLP模型训练中图片无法处理，暂时不放图片
		// patentOBJ.InstructionPic = FileToBase64(filepath.Dir(xmlPath))
	}
	patentOBJ.PatentType = pkg.PatentType[patentIndex]
	patentOBJ.XMLPath = xmlPath

	// 转换为JSON格式的字节数组
	jsonBytes, err := json.Marshal(patentOBJ)
	if err != nil {
		fmt.Println(err)
	}
	src := output + "/JSON" + "/" + fileName + ".json"
	// 将字节数组写入文件
	err = ioutil.WriteFile(src, jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	if NUM%MAXFILE == 0 {
		log.Println("parse done!", NUM)
		Hadoop.CreateDic(*client)
	}

	dst := filepath.Join(Hadoop.FileDic + "/" + fileName + ".json")
	//fmt.Println(src, dst)
	err = Hadoop.UploadFile(src, dst, *client)
	if err != nil {
		if strings.Contains(err.Error(), "file already exists") {
			ERRORNUN++
		} else {
			ERRORNUN++
			log.Println("UploadFile error:", err, src, ERRORNUN)
		}
	}

	NUM++
	// log.Println("parse done!", NUM)

	return nil
}

func retryDo(src, dst string, client hdfs.Client) {

	for {
		err := Hadoop.UploadFile(src, dst, client)
		if err == nil {
			return
		}
		if err != nil {
			fmt.Println("retry UploadFile", err)
		}
		time.Sleep(1000)
	}
}

// Par1Xml 主要解析 10-A 10-B 20-U 类型的专利
func Par1Xml(xmlPath, output string, patentIndex int, client *hdfs.Client) error {
	// log.Println("xml path -----", xmlPath, output, patentIndex)

	fileName := filepath.Base(xmlPath)
	fileName = strings.Split(fileName, ".")[0]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlPath); err != nil {
		log.Println(err, "解析失败手动处理========", xmlPath)
		return nil
	}
	root := doc.SelectElement("PatentDocumentAndRelated")
	if root == nil {
		log.Println("root is nil，解析失败手动处理========", xmlPath)
		return nil
	}

	resName := ""
	resNameE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/InventionTitle")
	if resNameE != nil {
		resName = resNameE.Text()
	}

	applicationCN := ""
	applicationCNE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/WIPOST3Code")
	if applicationCNE != nil {
		applicationCN = applicationCNE.Text()
	}
	applicationNO := ""
	applicationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/DocNumber")
	if applicationNOE != nil {
		applicationNO = applicationNOE.Text()
	}

	applicationDate := ""
	applicationDateE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/Date")
	if applicationDateE != nil {
		applicationDate = applicationDateE.Text()
	}

	publicationCN := ""
	publicationCNE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/WIPOST3Code")
	if publicationCNE != nil {
		publicationCN = publicationCNE.Text()
	}
	publicationNO := ""
	publicationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/DocNumber")
	if publicationNOE != nil {
		publicationNO = publicationNOE.Text()
	}

	publicationKind := ""
	publicationKindE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Kind")
	if publicationKindE != nil {
		publicationKind = publicationKindE.Text()
	}

	publicationDate := ""
	publicationDateE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Date")
	if publicationDateE != nil {
		publicationDate = publicationDateE.Text()
	}

	applicationAddress := ""
	applicationAddressE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Address/Text")
	if applicationAddressE != nil {
		applicationAddress = applicationAddressE.Text()
	}

	applicant := ""
	applicantE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Name")
	if applicantE != nil {
		applicant = applicantE.Text()
	}

	applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/BibliographicData/Parties/InventorDetails/Inventor")
	applicationInventors := ""
	if len(applicationInventorsList) != 0 {
		for _, v := range applicationInventorsList {
			g := v.FindElement("./AddressBook/Name").Text()
			applicationInventors = applicationInventors + g + ","
		}
	}

	abstract := ""
	abstractE := doc.FindElement("./PatentDocumentAndRelated/Abstract/Paragraphs")
	if abstractE != nil {
		abstract = abstractE.Text()
	}

	claimList := doc.FindElements("./PatentDocumentAndRelated/Claims/Claim")
	claim := ""
	if len(claimList) != 0 {
		for _, v := range claimList {
			for _, v2 := range v.ChildElements() {
				g := v2.Text()
				claim = claim + g
			}
		}
	}

	mainClassificationNO := ""
	mainClassificationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR/Text")
	if mainClassificationNOE != nil {
		mainClassificationNO = mainClassificationNOE.Text()
	}

	ClassificationNO := ""
	ClassificationNOList := doc.FindElements("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR")
	if len(ClassificationNOList) != 0 {
		for _, v := range ClassificationNOList {
			g := v.SelectElement("Text").Text()
			ClassificationNO = ClassificationNO + g + ","
		}
	}

	Agency := ""
	Agent := ""
	agency := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties")
	if agency != nil {
		for _, v := range agency.ChildElements() {
			if v.Tag == "AgentDetails" {
				g := v.FindElement("./Agent/Agency/AddressBook/OrganizationName").Text()
				f := v.FindElement("./Agent/AddressBook/Name").Text()
				Agency = g
				Agent = f
			}
		}
	}

	// 一般专利描述里把 技术领域、背景技术、发明内容、附图说明、具体实施方式 都融合在一起，需要将这些内容提取出来

	technicalField := ""
	technicalBackground := ""
	context := ""
	instructionWithPicture := ""
	implementation := ""
	descriptionE := doc.FindElement("./PatentDocumentAndRelated/Description")
	// 单独字段,没有子字段
	dE := doc.FindElement("./PatentDocumentAndRelated/Description/TechnicalField")

	// 融合字段
	describe := doc.FindElements("./PatentDocumentAndRelated/Description/Paragraphs")

	// 单独字段，但是有子字段
	// 技术领域
	tfE := doc.FindElements("./PatentDocumentAndRelated/Description/TechnicalField/Paragraphs")
	// 技术背景
	baE := doc.FindElements("./PatentDocumentAndRelated/Description/BackgroundArt/Paragraphs")
	// 发明内容
	disE := doc.FindElements("./PatentDocumentAndRelated/Description/Disclosure/Paragraphs")
	// 附图说明
	ddE := doc.FindElements("./PatentDocumentAndRelated/Description/DrawingsDescription/Paragraphs")
	// 具体实施方式
	imE := doc.FindElements("./PatentDocumentAndRelated/Description/InventionMode/Paragraphs")

	if descriptionE != nil {
		if len(tfE) == 0 && len(baE) == 0 && len(disE) == 0 && len(ddE) == 0 && len(imE) == 0 {
			if dE != nil {
				technicalField = doc.FindElement("./PatentDocumentAndRelated/Description/TechnicalField").Text()
				technicalBackground = doc.FindElement("./PatentDocumentAndRelated/Description/BackgroundArt").Text()
				context = doc.FindElement("./PatentDocumentAndRelated/Description/Disclosure").Text()
				instructionWithPicture = doc.FindElement("./PatentDocumentAndRelated/Description/DrawingsDescription").Text()
				implementation = doc.FindElement("./PatentDocumentAndRelated/Description/InventionMode").Text()
			} else {

				var DescribeArr []string
				for _, v := range describe {
					DescribeArr = append(DescribeArr, v.Text())
				}
				technicalField, technicalBackground, context, instructionWithPicture, implementation = ParDescribeArr(DescribeArr, patentIndex)
			}

		} else {
			// 该专利 有单独的技术领域、技术背景等字段
			// 技术领域
			if len(tfE) != 0 {
				var tfEArr []string
				for _, v := range tfE {
					g := v.Text()
					tfEArr = append(tfEArr, g)
				}
				technicalField = strings.Join(tfEArr, ",")
			}
			// 技术背景
			if len(baE) != 0 {
				var baEArr []string
				for _, v := range baE {
					g := v.Text()
					baEArr = append(baEArr, g)
				}
				technicalBackground = strings.Join(baEArr, ",")
			}
			// 发明内容
			if len(disE) != 0 {
				var disEArr []string
				for _, v := range disE {
					g := v.Text()
					disEArr = append(disEArr, g)
				}
				context = strings.Join(disEArr, ",")
			}

			// 具体实施方式
			if len(imE) != 0 {
				var imEArr []string

				for _, v := range imE {
					g := v.Text()
					imEArr = append(imEArr, g)
				}
				implementation = strings.Join(imEArr, ",")
			}
			// 附图说明
			if len(ddE) != 0 {
				var ddEArr []string
				for _, v := range ddE {
					g := v.Text()
					ddEArr = append(ddEArr, g)
				}
				instructionWithPicture = strings.Join(ddEArr, ",")
			}

		}
	}

	patentObj := pkg.Patent{
		Name:                   resName,
		ApplicationNO:          applicationCN + applicationNO,
		ApplicationDate:        applicationDate,
		PublicationNO:          publicationCN + publicationNO + publicationKind,
		PublicationDate:        publicationDate,
		Applicant:              applicant,
		ApplicantAddress:       applicationAddress,
		Inventors:              applicationInventors,
		Abstract:               abstract,             //摘要
		Claim:                  claim,                //主权项
		MainClassificationNO:   mainClassificationNO, //主分类号
		ClassificationNO:       ClassificationNO,     //分类号
		Agency:                 Agency,
		Agent:                  Agent,
		PatentType:             pkg.PatentType[patentIndex],
		TechnicalField:         technicalField,
		TechnicalBackground:    technicalBackground,
		Context:                context,
		InstructionWithPicture: instructionWithPicture,
		Implementation:         implementation,
		//InstructionPic: FileToBase64(filepath.Dir(xmlPath)), NLP模型训练中图片无法处理，暂时不放图片
		AbstractPic: "",
		XMLPath:     xmlPath,
	}

	// 转换为JSON格式的字节数组
	jsonBytes, err := json.Marshal(patentObj)
	if err != nil {
		fmt.Println(err)
	}
	src := output + "/JSON" + "/" + fileName + ".json"
	// 将字节数组写入文件
	err = ioutil.WriteFile(src, jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	if NUM%MAXFILE == 0 {
		log.Println("parse done!", NUM)
		Hadoop.CreateDic(*client)
	}

	dst := filepath.Join(Hadoop.FileDic + "/" + fileName + ".json")
	// fmt.Println(src, dst)
	err = Hadoop.UploadFile(src, dst, *client)
	if err != nil {
		if strings.Contains(err.Error(), "file already exists") {
			ERRORNUN++
		} else {
			ERRORNUN++
			log.Println("UploadFile error:", err, src, ERRORNUN)
		}
	}

	NUM++
	// log.Println("parse done!", NUM)

	return nil
}

// ParDescribeArr 把描述都融合在一起
func ParDescribeArr(DescribeArr []string, patentIndex int) (string1, string2, string3, string4, string5 string) {
	index1 := -1 // 技术领域
	index2 := -1 // 背景"技术
	index3 := -1 // 发明内容
	index4 := -1 // 附图说明
	index5 := -1 // 具体实施方式
	for i := 0; i < len(DescribeArr); i++ {
		if strings.Contains(DescribeArr[i], "技术领域") {
			if index1 == -1 {
				index1 = i
			}
		}

		if strings.Contains(DescribeArr[i], "背景") && strings.Contains(DescribeArr[i], "技术") {
			if index2 == -1 {
				index2 = i
			}
		}

		if strings.Contains(DescribeArr[i], "发明内容") {
			if index3 == -1 {
				index3 = i
			}
		}

		if strings.Contains(DescribeArr[i], "附图说明") {
			if index4 == -1 {
				index4 = i
			}
		}

		if strings.Contains(DescribeArr[i], "具体实施方式") {
			if index5 == -1 {
				index5 = i
			}
		}

	}

	// fmt.Println(patentIndex, index1, index2, index3, index4, index5)

	var stringArr1, stringArr2, stringArr3, stringArr4, stringArr5 []string
	if index1 != -1 && index2 != -1 && index3 != -1 && index4 != -1 && index5 != -1 {
		indexArr := []int{index1, index2, index3, index4, index5}
		sort.Ints(indexArr)
		for k, v := range indexArr {
			if v == index1 {
				if k != 4 {
					stringArr1 = DescribeArr[v:indexArr[k+1]]
				} else {
					stringArr1 = DescribeArr[v:]
				}
			}
			if v == index2 {
				if k != 4 {
					stringArr2 = DescribeArr[v:indexArr[k+1]]
				} else {
					stringArr2 = DescribeArr[v:]
				}
			}
			if v == index3 {
				if k != 4 {
					stringArr3 = DescribeArr[v:indexArr[k+1]]
				} else {
					stringArr3 = DescribeArr[v:]
				}
			}
			if v == index4 {
				if k != 4 {
					stringArr4 = DescribeArr[v:indexArr[k+1]]
				} else {
					stringArr4 = DescribeArr[v:]
				}
			}
			if v == index5 {
				if k != 4 {
					stringArr5 = DescribeArr[v:indexArr[k+1]]
				} else {
					stringArr5 = DescribeArr[v:]
				}
			}
		}
	} else {
		return "", "", CombineStr(DescribeArr), "", ""
	}

	return CombineStr(stringArr1), CombineStr(stringArr2), CombineStr(stringArr3), CombineStr(stringArr4), CombineStr(stringArr5)
}

func CombineStr(arr []string) string {
	tmp := ""
	for _, v := range arr {
		tmp = tmp + v
	}
	return tmp
}
