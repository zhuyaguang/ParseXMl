package Hadoop

import (
	"testing"
)

func TestConnectHadoop(t *testing.T) {
	clinet := ConnectHadoop()
	UploadFile("test2.json", "/test2/test3.json", *clinet)
}
