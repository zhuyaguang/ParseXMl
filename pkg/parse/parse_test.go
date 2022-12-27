package parse

import (
	zgorm "patentExtr/pkg/gorm"
	"testing"
)

func TestPar1Xml(t *testing.T) {
	EngineMysqlGORM := zgorm.ConnectMysql()
	//EngineMysqlGORM.AutoMigrate(&pkg.Patent{})
	Par1Xml("test2.XML", "./", 1, EngineMysqlGORM)
}
