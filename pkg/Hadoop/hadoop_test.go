package Hadoop

import (
	"log"
	"path/filepath"
	"testing"
)

func TestConnectHadoop(t *testing.T) {
	clinet := ConnectHadoop()
	dst := filepath.Join(FileDic + "/" + "test2" + ".json")
	err := UploadFile("test2.json", dst, *clinet)
	if err != nil {
		log.Fatal(err)
	}
}
git