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
	XMLPath                string
}

func (Patent) TableName() string {
	return "patent-20230322"
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
type Document struct {
	Timestamp      string `json:"@timestamp"`
	Version        string `json:"@version"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
	ID             string `json:"id"`
	AuthorizedDate string `json:"authorized_data"`
	AuthorizedNO   string `json:"authorized_no"`
	AreaCode       string `json:"area_code"` // 国省代码
	InventorsID    string `json:"inventors_id"`
	InventorsName  string `json:"inventors_name"`
	InventorsUnit  string `json:"inventors_unit"`
	KeyWords       string `json:"keywords"`
	LegalStatus    string `json:"legal_status"`
	Subject        string `json:"subject"`
	UnitType       string `json:"unit_type"`
	URL            string `json:"url"`
	Year           string `json:"year"`
	Tags           string `json:"tags"`

	Name                 string `json:"name"`
	Abstract             string `json:"abstract_ch"`            //摘要
	Agency               string `json:"agency"`                 // 代理机构
	Agent                string `json:"agent"`                  // 代理人
	Applicant            string `json:"applicant"`              //申请人
	ApplicantAddress     string `json:"applicant_address"`      //申请地址
	ApplicationDate      string `json:"application_date"`       //申请日期
	ApplicationNO        string `json:"application_no"`         //申请号
	Claim                string `json:"claim"`                  //权利要求书
	MainClassificationNO string `json:"main_classification_no"` //主分类号
	ClassificationNO     string `json:"classification_no"`      //分类号
	Inventors            string `json:"inventors"`              //发明人
	PatentType           string `json:"patent_type"`            // 专利类型
	PublicationNO        string `json:"publication_no"`         // 公开号
	PublicationDate      string `json:"publication_date"`       // 公开日期

	TechnicalField         string `json:"technical_field"`          // 技术领域
	TechnicalBackground    string `json:"technical_background"`     // 背景技术
	Context                string `json:"context"`                  // 发明内容
	InstructionWithPicture string `json:"instruction_with_picture"` // 附图说明
	Implementation         string `json:"implementation"`           // 具体实施方法
	InstructionPic         string `json:"instruction_pic"`          // 说明书附图
	AbstractPic            string `json:"abstract_pic"`             // 摘要附图
	XMLPath                string `json:"xml_path"`                 // 原 xml 文件地址
}

type ESPatent struct {
	Name                   string
	ApplicationNO          string //申请号
	ApplicationDate        string //申请日期
	PublicationNO          string `gorm:"index:idx_public_no,unique"` // 公开号
	PublicationDate        string // 公开日期
	Applicant              string //申请人
	ApplicantAddress       string //申请地址
	Inventors              string //发明人
	Abstract               string //摘要
	Claim                  string //权利要求书
	MainClassificationNO   string //主分类号
	ClassificationNO       string //分类号
	Agency                 string // 代理机构
	Agent                  string // 代理人
	PatentType             string // 专利类型
	TechnicalField         string // 技术领域
	TechnicalBackground    string // 背景技术
	Context                string // 发明内容
	InstructionWithPicture string // 附图说明
	Implementation         string // 具体实施方法
	InstructionPic         string // 说明书附图
	AbstractPic            string // 摘要附图
	XMLPath                string // 原 xml 文件地址
	AreaCode               string // 国省代码
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
