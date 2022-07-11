package parse

import (
	"encoding/json"
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"patentExtr/pkg"
	"path/filepath"
	"strings"
)

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

	resName :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/DesignTitle").Text()

	applicationNO := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PatentNumber").Text()
	applicationDate :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicationReference/DocumentID/Date").Text()

	publicationCN := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/WIPOST3Code").Text()
	publicationNO := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/DocNumber").Text()
	publicationKind := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Kind").Text()
	publicationDate :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/PublicationReference/DocumentID/Date").Text()

	applicationAddress :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/CorrespondenceAddress/AddressBook/Address/Text").Text()

	applicant :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ApplicantDetails/Applicant/AddressBook/Name").Text()

	applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/DesignBibliographicData/DesignerDetails/Designer")
	applicationInventors:=""
	for _,v:=range applicationInventorsList{
		g:=v.FindElement("./AddressBook/Name").Text()
		applicationInventors=applicationInventors+g+","
	}


	abstractList :=doc.FindElements("./PatentDocumentAndRelated/DesignBriefExplanation/Paragraphs")
	abstract :=""
	for _,v:=range abstractList{
		g:=v.Text()
		abstract=abstract+g
	}

	mainClassificationNO := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/ClassificationLocarno/MainClassification").Text()

	Agency := doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/Agency/AddressBook/OrganizationName").Text()
	Agent :=doc.FindElement("./PatentDocumentAndRelated/DesignBibliographicData/AgentDetails/Agent/AddressBook/Name").Text()


	patentObj:=pkg.Patent{
		Name:                   resName,
		ApplicationNO:          applicationNO,
		ApplicationDate:        applicationDate,
		PublicationNO:          publicationCN+publicationNO+publicationKind,
		PublicationDate:        publicationDate,
		Applicant:              applicant,
		ApplicantAddress:       applicationAddress,
		Inventors:              applicationInventors,
		Abstract:               abstract,
		Claim:                  "",
		MainClassificationNO: mainClassificationNO,
		ClassificationNO:       mainClassificationNO,
		Agency:                 Agency,
		Agent:                  Agent,
		PatentType:             pkg.PatentType[patentIndex],
		InstructionPic: "",
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

	resName :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/InventionTitle").Text()


	applicationCN := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/WIPOST3Code").Text()
	applicationNO := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/DocNumber").Text()
	applicationDate :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ApplicationReference/DocumentID/Date").Text()


	publicationCN := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/WIPOST3Code").Text()
	publicationNO := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/DocNumber").Text()
	publicationKind := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Kind").Text()
	publicationDate :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/PublicationReference/DocumentID/Date").Text()


	applicationAddress :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Address/Text").Text()
	applicant :=doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties/ApplicantDetails/Applicant/AddressBook/Name").Text()

	applicationInventorsList := doc.FindElements("./PatentDocumentAndRelated/BibliographicData/Parties/InventorDetails/Inventor")
	applicationInventors:=""
	for _,v:=range applicationInventorsList{
		g:=v.FindElement("./AddressBook/Name").Text()
		applicationInventors=applicationInventors+g+","
	}

	abstract :=doc.FindElement("./PatentDocumentAndRelated/Abstract/Paragraphs").Text()

	claimList := doc.FindElements("./PatentDocumentAndRelated/Claims/Claim")
	claim :=""
	for _,v:=range claimList{
		g:=v.SelectElement("ClaimText").Text()
		claim = claim+g
	}

	mainClassificationNO := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR/Text").Text()

	ClassificationNO:=""
	ClassificationNOList :=doc.FindElements("./PatentDocumentAndRelated/BibliographicData/ClassificationIPCRDetails/ClassificationIPCR")
	for _,v :=range ClassificationNOList{
		g :=v.SelectElement("Text").Text()
		ClassificationNO = ClassificationNO+g+","
	}

	agency := doc.FindElement("./PatentDocumentAndRelated/BibliographicData/Parties")
	Agency :=""
	Agent :=""
	for _,v :=range agency.ChildElements(){
		if v.Tag == "AgentDetails"{
			g:=v.FindElement("./Agent/Agency/AddressBook/OrganizationName").Text()
			f:=v.FindElement("./Agent/AddressBook/Name").Text()
			Agency= g
			Agent = f
		}
	}


	describe := doc.FindElements("./PatentDocumentAndRelated/Description/Paragraphs")
	var DescribeArr []string
	for _,v:=range describe{
		DescribeArr = append(DescribeArr,v.Text())
	}

	technicalField,technicalBackground,context,instructionWithPicture,implementation:=ParDescribeArr(DescribeArr,patentIndex)



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
		InstructionPic: "",
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

		stringArr1 =DescribeArr[index1:index2]
		stringArr2 =DescribeArr[index2:index3]
		stringArr3 =DescribeArr[index3:index4]
		stringArr4 =DescribeArr[index4:index5]
		stringArr5 =DescribeArr[index5:]



	return CombineStr(stringArr1),CombineStr(stringArr2),CombineStr(stringArr3),CombineStr(stringArr4),CombineStr(stringArr5)
}

func CombineStr(arr []string)string  {
	tmp:=""
	for _,v:=range arr{
		tmp=tmp+v
	}
	return tmp
}