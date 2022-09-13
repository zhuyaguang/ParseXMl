package pkg

import "gorm.io/gorm"

const PATENT1 = "中国外观设计专利授权公告标准化著录项目及切图数据" //30-S
const PATENT2 = "中国发明专利申请公布标准化全文文本数据"      //10-A
const PATENT3 = "中国发明专利授权公告标准化全文文本数据"      //10-B
const PATENT4 = "中国实用新型专利授权公告标准化全文文本数据"    //20-U

var PatentType = [4]string{"中国外观设计专利授权公告标准化著录项目及切图数据", "中国发明专利申请公布标准化全文文本数据", "中国发明专利授权公告标准化全文文本数据", "中国实用新型专利授权公告标准化全文文本数据"}

type Patent struct {
	gorm.Model
	Name                   string
	ApplicationNO          string
	ApplicationDate        string
	PublicationNO          string `gorm:"index:idx_public_no,unique"`
	PublicationDate        string
	Applicant              string
	ApplicantAddress       string
	Inventors              string
	Abstract               string
	Claim                  string
	MainClassificationNO   string
	ClassificationNO       string
	Agency                 string
	Agent                  string
	PatentType             string
	TechnicalField         string
	TechnicalBackground    string
	Context                string
	InstructionWithPicture string
	Implementation         string
	InstructionPic         string
	AbstractPic            string
}

type mysqlPatent struct {
	gorm.Model
	Name                   string `json:"name" json:"name,omitempty"`
	ApplicationNO          string `json:"application_no" json:"application_no,omitempty"`
	ApplicationDate        string `json:"application_date,omitempty"`
	PublicationNO          string `json:"publication_no,omitempty"`
	PublicationDate        string `json:"publication_date,omitempty"`
	Applicant              string `json:"applicant,omitempty"`
	ApplicantAddress       string `json:"applicant_address,omitempty"`
	Inventors              string `json:"inventors,omitempty"`
	Abstract               string `json:"abstract,omitempty"`
	Claim                  string `json:"claim,omitempty"`
	MainClassificationNO   string `json:"main_classification_no,omitempty"`
	ClassificationNO       string `json:"classification_no,omitempty"`
	Agency                 string `json:"agency,omitempty"`
	Agent                  string `json:"agent,omitempty"`
	PatentType             string `json:"patent_type,omitempty"`
	TechnicalField         string `json:"technical_field,omitempty"`
	TechnicalBackground    string `json:"technical_background,omitempty"`
	Context                string `json:"context,omitempty"`
	InstructionWithPicture string `json:"instruction_with_picture,omitempty"`
	Implementation         string `json:"implementation,omitempty"`
	InstructionPic         string `json:"instruction_pic,omitempty"`
	AbstractPic            string `json:"abstract_pic,omitempty"`
}

//`name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '名称',
//`application_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '申请号',
//`application_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '申请日期',
//`publication_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '公布号',
//`publication_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '公布日期',
//`applicant` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '申请人',
//`applicant_address` varchar(1024) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '申请人地址',
//`inventors` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT '发明人',
//`abstract_ch` longtext CHARACTER SET utf8 COLLATE utf8_bin COMMENT '摘要',
//`claim` longtext CHARACTER SET utf8 COLLATE utf8_bin COMMENT '主权项',
//`main_classification_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '主分类号',
//`classification_no` varchar(1024) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '分类号',
//`agency` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '代理机构',
//`agent` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '代理人',
//`patent_type` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '专利类别,
//`technical_field` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '技术领域,
//`technical_background` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '背景技术,
//`content` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '发明内容,
//`instruction_with_picture` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '附图说明,
//`implementation` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '具体实施方式,
//`instruction_pic` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '说明书附图,
//`abstract_pic` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '摘要附图,
