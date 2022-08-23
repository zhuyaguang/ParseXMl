package parse

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"log"
	"net/http"
	"patentExtr/pkg"
	"path/filepath"
	"strings"
)

var NUM = 0

func Par0Xml(xmlPath ,output string,patentIndex int) error {
	fmt.Println("xml path -----",xmlPath,output,patentIndex)

	fileName :=filepath.Base(xmlPath)

	fileName = strings.Split(fileName,".")[0]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlPath); err != nil {
		panic(err)
	}
	root := doc.SelectElement("PatentDocumentAndRelated")
	fmt.Println("ROOT element:", root.Tag)

	patentOBJ :=pkg.Patent{}

	resNameE :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/DesignTitle")
	if resNameE !=nil{
		patentOBJ.Name=resNameE.Text()
		patentOBJ.ApplicationNO = doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PatentNumber").Text()
		patentOBJ.ApplicationDate=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicationReference/DocumentID/Date").Text()

		publicationCN := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/WIPOST3Code").Text()
		publicationNO := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/DocNumber").Text()
		publicationKind := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Kind").Text()
		patentOBJ.PublicationNO = publicationCN+publicationNO+publicationKind
		patentOBJ.PublicationDate =doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Date").Text()


		applicationAddressE :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/CorrespondenceAddress/AddressBook/Address/Text")
		if applicationAddressE == nil{
			applicationAddressE =doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicantDetails/Applicant/AddressBook/Address/Text")
		}
		patentOBJ.ApplicantAddress= applicationAddressE.Text()


		patentOBJ.Applicant=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicantDetails/Applicant/AddressBook/Name").Text()

		applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/DesignBibliographicData/DesignerDetails/Designer")
		applicationInventors:=""
		for _,v:=range applicationInventorsList{
			g:=v.FindElement("./AddressBook/Name").Text()
			applicationInventors=applicationInventors+g+","
		}
		patentOBJ.Inventors=applicationInventors


		abstractList :=doc.FindElements("./PatentDocumentAndRelated/DesignBriefExplanation/Paragraphs")
		abstract :=""
		for _,v:=range abstractList{
			g:=v.Text()
			abstract=abstract+g
		}
		patentOBJ.Abstract = abstract

		patentOBJ.MainClassificationNO= doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ClassificationLocarno/MainClassification").Text()

		AgencyE := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/Agency/AddressBook/OrganizationName")
		if AgencyE !=nil{
			patentOBJ.Agency = AgencyE.Text()
		}


		AgentE :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/AddressBook/Name")
		if AgentE !=nil{
			patentOBJ.Agent = AgentE.Text()
		}

		// NLP模型训练中图片无法处理，暂时不放图片
		// patentOBJ.InstructionPic = FileToBase64(filepath.Dir(xmlPath))
	}

	file, err := json.MarshalIndent(patentOBJ, "", " ")
	if err!=nil{
		return err
	}

	err = ioutil.WriteFile(output+"/"+fileName+".json", file, 0644)
	if err!=nil{
		return err
	}
	NUM++
	fmt.Println("parse done!",NUM)

	return nil
}


func Par1Xml(xmlPath ,output string,patentIndex int) error {
	fmt.Println("xml path -----",xmlPath,output,patentIndex)
	fileName :=filepath.Base(xmlPath)

	fileName = strings.Split(fileName,".")[0]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlPath); err != nil {
		panic(err)
	}
	root := doc.SelectElement("PatentDocumentAndRelated")
	fmt.Println("ROOT element:", root.Tag)


	resName:=""
	resNameE :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/InventionTitle")
	if resNameE !=nil{
		resName=resNameE.Text()
	}

	applicationCN := ""
	applicationCNE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/WIPOST3Code")
	if applicationCNE !=nil{
		applicationCN = applicationCNE.Text()
	}
	applicationNO:=""
	applicationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/DocNumber")
	if applicationNOE !=nil{
		applicationNO=applicationNOE.Text()
	}

	applicationDate:=""
	applicationDateE :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/Date")
	if applicationDateE!=nil{
		applicationDate = applicationDateE.Text()
	}

	publicationCN:=""
	publicationCNE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/WIPOST3Code")
	if publicationCNE!=nil{
		publicationCN=publicationCNE.Text()
	}
	publicationNO :=""
	publicationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/DocNumber")
	if publicationNOE !=nil{
		publicationNO= publicationNOE.Text()
	}

	publicationKind :=""
	publicationKindE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Kind")
	if publicationKindE !=nil{
		publicationKind = publicationKindE.Text()
	}

	publicationDate :=""
	publicationDateE :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Date")
	if publicationDateE !=nil{
		publicationDate = publicationDateE.Text()
	}

	applicationAddress := ""
	applicationAddressE :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Address/Text")
	if applicationAddressE !=nil{
		applicationAddress = applicationAddressE.Text()
	}

	applicant :=""
	applicantE :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Name")
	if applicantE !=nil{
		applicant = applicantE.Text()
	}

	applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/BibliographicData/Parties/InventorDetails/Inventor")
	applicationInventors:=""
	if len(applicationInventorsList) !=0{
		for _,v:=range applicationInventorsList{
			g:=v.FindElement("./AddressBook/Name").Text()
			applicationInventors=applicationInventors+g+","
		}
	}

	abstract :=""
	abstractE :=doc.FindElement("./PatentDocumentAndRelated/Abstract/Paragraphs")
	if abstractE !=nil{
		abstract = abstractE.Text()
	}

	claimList := doc.FindElements("./PatentDocumentAndRelated/Claims/Claim")
	claim :=""
	if len(claimList) !=0{
		for _,v:=range claimList{
			g:=v.SelectElement("ClaimText").Text()
			claim = claim+g
		}
	}

	mainClassificationNO :=""
	mainClassificationNOE := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR/Text")
	if mainClassificationNOE!=nil{
		mainClassificationNO = mainClassificationNOE.Text()
	}

	ClassificationNO:=""
	ClassificationNOList :=doc.FindElements("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR")
	if len(ClassificationNOList) !=0{
		for _,v :=range ClassificationNOList{
			g :=v.SelectElement("Text").Text()
			ClassificationNO = ClassificationNO+g+","
		}
	}

	Agency :=""
	Agent :=""
	agency := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties")
	if agency !=nil{
		for _,v :=range agency.ChildElements(){
			if v.Tag == "AgentDetails"{
				g:=v.FindElement("./Agent/Agency/AddressBook/OrganizationName").Text()
				f:=v.FindElement("./Agent/AddressBook/Name").Text()
				Agency= g
				Agent = f
			}
		}
	}



	technicalField:=""
	technicalBackground:=""
	context:=""
	instructionWithPicture:=""
	implementation:=""
	descriptionE :=doc.FindElement("./PatentDocumentAndRelated/Description")
	dE := doc.FindElement("./PatentDocumentAndRelated/Description/TechnicalField")
	describe := doc.FindElements("./PatentDocumentAndRelated/Description/Paragraphs")

	tfE := doc.FindElements("./PatentDocumentAndRelated/Description/TechnicalField/Paragraphs")
	baE :=doc.FindElements("./PatentDocumentAndRelated/Description/BackgroundArt/Paragraphs")
	disE :=doc.FindElements("./PatentDocumentAndRelated/Description/Disclosure/Paragraphs")
	ddE :=doc.FindElements("./PatentDocumentAndRelated/Description/DrawingsDescription/Paragraphs")
	imE :=doc.FindElements("./PatentDocumentAndRelated/Description/InventionMode/Paragraphs")


	if descriptionE !=nil{
		if len(tfE)==0 && len(baE)==0 && len(disE)==0 && len(ddE)==0 && len(imE)==0{
			if dE !=nil{
				technicalField = doc.FindElement("./PatentDocumentAndRelated/Description/TechnicalField").Text()
				technicalBackground = doc.FindElement("./PatentDocumentAndRelated/Description/BackgroundArt").Text()
				context = doc.FindElement("./PatentDocumentAndRelated/Description/Disclosure").Text()
				instructionWithPicture = doc.FindElement("./PatentDocumentAndRelated/Description/DrawingsDescription").Text()
				implementation = doc.FindElement("./PatentDocumentAndRelated/Description/InventionMode").Text()
			}else{

				var DescribeArr []string
				for _,v:=range describe{
					DescribeArr = append(DescribeArr,v.Text())
				}
				technicalField,technicalBackground,context,instructionWithPicture,implementation =ParDescribeArr(DescribeArr,patentIndex)
			}

		}else{
			if len(tfE) !=0{
				for _,v :=range tfE{
					g :=v.Text()
					technicalField = technicalField+g
				}
			}

			if len(baE) !=0{
				for _,v :=range baE{
					g :=v.Text()
					technicalBackground = technicalBackground+g
				}
			}
		}
	}



	patentObj:=pkg.Patent{
		Name:                   resName,
		ApplicationNO:          applicationCN+applicationNO,
		ApplicationDate:        applicationDate,
		PublicationNO:          publicationCN+publicationNO+publicationKind,
		PublicationDate:        publicationDate,
		Applicant:              applicant,
		ApplicantAddress:       applicationAddress,
		Inventors:              applicationInventors,
		Abstract:               abstract,
		Claim:                  claim,
		MainClassificationNO: mainClassificationNO,
		ClassificationNO:       ClassificationNO,
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
	}
	file, err := json.MarshalIndent(patentObj, "", " ")
	if err!=nil{
		return err
	}

	err = ioutil.WriteFile(output+"/"+fileName+".json", file, 0644)
	if err!=nil{
		return err
	}
	NUM++
	fmt.Println("parse done!",NUM)

	return nil
}

func ParDescribeArr(DescribeArr []string,patentIndex int) (string1,string2,string3,string4,string5 string) {
	index1:=-1  // 技术领域
	index2:=-1  // 背景"技术
	index3:=-1  // 发明内容
	index4:=-1  // 附图说明
	index5:=-1  // 具体实施方式
	for i:=0;i<len(DescribeArr);i++{
		if strings.Contains(DescribeArr[i],"技术领域"){
			if index1 == -1{
				index1 = i
			}
		}

		if strings.Contains(DescribeArr[i],"背景")&&strings.Contains(DescribeArr[i],"技术"){
			if index2 == -1{
				index2 = i
			}
		}

		if strings.Contains(DescribeArr[i],"发明内容"){
			if index3 == -1{
				index3 = i
			}
		}

		if strings.Contains(DescribeArr[i],"附图说明"){
			if index4 == -1{
				index4 = i
			}
		}

		if strings.Contains(DescribeArr[i],"具体实施方式"){
			if index5 == -1{
				index5 = i
			}
		}

	}

	fmt.Println(patentIndex,index1,index2,index3,index4,index5)


	var stringArr1,stringArr2,stringArr3,stringArr4,stringArr5 []string
	if index1!=-1&&index2!=-1&&index3!=-1&&index4!=-1&&index5!=-1{
		stringArr1 =DescribeArr[index1:index2]
		stringArr2 =DescribeArr[index2:index3]
		stringArr3 =DescribeArr[index3:index4]
		stringArr4 =DescribeArr[index4:index5]
		stringArr5 =DescribeArr[index5:]
	}else{
		return "","",CombineStr(DescribeArr),"",""
	}




	return CombineStr(stringArr1),CombineStr(stringArr2),CombineStr(stringArr3),CombineStr(stringArr4),CombineStr(stringArr5)
}

func CombineStr(arr []string)string  {
	tmp:=""
	for _,v:=range arr{
		tmp=tmp+v
	}
	return tmp
}

func FileToBase64(filePath string) string  {
	fmt.Println("=======",filePath)
	base64String :=""
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(),"JPG") || strings.HasSuffix(f.Name(),"TIF") {
			picFile :=filePath+"/"+f.Name()
			base64String = base64String+ConvertToBase64(picFile)+","
		}
	}
	return base64String

}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ConvertToBase64(filePath string) string{
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	//fmt.Println(base64Encoding)

	return base64Encoding
}