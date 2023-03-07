package Hadoop

import (
	"github.com/colinmarc/hdfs"
	"log"
	"path/filepath"
	"time"
)

const FOOTPATH = "/patentJson"

var FileDic string

func ConnectHadoop() *hdfs.Client {
	client, err := hdfs.New("10.100.29.40:8020")
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	// 将日期格式化为指定格式
	dateStr := now.Format("200601021504")
	// 创建文件夹
	FileDic = filepath.Join(FOOTPATH + "/" + dateStr)
	err = client.Mkdir(FileDic, 0777)
	if err != nil {
		log.Println(err)
	}
	return client
}

func UploadFile(src, dst string, client hdfs.Client) error {

	// 上传文件

	err := client.CopyToRemote(src, dst)
	if err != nil {
		log.Println(err)
		return err
	}

	// 列出文件夹内容
	//infos, err := client.ReadDir(FOOTPATH)
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	//
	//for _, info := range infos {
	//	fmt.Println(info.Name())
	//}
	return nil
}
