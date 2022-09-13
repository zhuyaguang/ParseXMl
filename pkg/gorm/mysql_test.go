package gorm

import (
	"patentExtr/pkg"
	"testing"
)

func TestCreate(t *testing.T) {
	db :=ConnectMysql()
	p :=pkg.Patent{
		Name:                   "123",
		ApplicationNO:          "",
		ApplicationDate:        "",
		PublicationNO:          "",
		PublicationDate:        "",
		Applicant:              "",
		ApplicantAddress:       "",
		Inventors:              "",
		Abstract:               "",
		Claim:                  "",
		MainClassificationNO:   "",
		ClassificationNO:       "",
		Agency:                 "",
		Agent:                  "",
		PatentType:             "",
		TechnicalField:         "",
		TechnicalBackground:    "",
		Context:                "",
		InstructionWithPicture: "",
		Implementation:         "",
		InstructionPic:         "",
		AbstractPic:            "",
	}
	Create(p,db)
	//Search(db)
}
