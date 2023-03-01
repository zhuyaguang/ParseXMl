package parse

import (
	"patentExtr/pkg/Hadoop"
	"testing"
)

func TestPar1Xml(t *testing.T) {
	EngineHadoopGORM := Hadoop.ConnectHadoop()
	//EngineMysqlGORM.AutoMigrate(&pkg.Patent{})
	Par1Xml("test2.XML", "./", 1, EngineHadoopGORM)
}
