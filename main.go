package main

import (
	"flag"
	"fmt"
	"github.com/colinmarc/hdfs"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"patentExtr/pkg"
	"patentExtr/pkg/Hadoop"
	"patentExtr/pkg/parse"
	"path/filepath"
	"strings"
	"time"
)

var endTime = [4]string{"", "", "", ""}
var EngineMysqlGORM *gorm.DB
var EngineHadoopGORM *hdfs.Client

func init() {
	//EngineMysqlGORM = zgorm.ConnectMysql()
	//EngineMysqlGORM.AutoMigrate(&pkg.Patent{})
	EngineHadoopGORM = Hadoop.ConnectHadoop()
}

func main() {
	dataAdd := flag.String("data", "/data/sipo", "source data address")
	outputAdd := flag.String("output", "/data/output", "output xml address")
	SStart := flag.String("s-start", "20220101", "30-s start parse time")
	AStart := flag.String("a-start", "20220101", "10-a start parse time")
	BStart := flag.String("b-start", "20220101", "10-b start parse time")
	UStart := flag.String("u-start", "20220101", "20-u start parse time")
	flag.Parse()

	log.Println(*dataAdd, *outputAdd, *SStart, *AStart, *BStart, *UStart)
	endTime = [4]string{*SStart, *AStart, *BStart, *UStart}

	start := time.Now()
	// Code to measure
	duration := time.Since(start)

	// 把专利数据解压到 output 目录
	// extractingXml(*dataAdd, *outputAdd)

	log.Println(duration)

	// 解析 XML 数据成 json 文件
	err := findXML(*outputAdd)
	if err != nil {
		log.Println(err.Error())
	}

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	log.Println(duration)

}

func extractingXml(dirPath string, output string) error {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// 办理第一层 文件夹目录含有 IMGS-30-S 或者 TXTS
		if strings.Contains(f.Name(), "TXTS") || strings.Contains(f.Name(), "IMGS-30-S") {
			patentType := f.Name()
			log.Println("file-type", patentType)
			// 进入 专利 目录
			if f.IsDir() {
				subfiles, err := ioutil.ReadDir(dirPath + "/" + patentType)
				if err != nil {
					log.Fatal(err)
				}
				for _, f := range subfiles {
					// 进入 日期 目录
					if f.IsDir() {
						patentdir := f.Name()
						//log.Println("file-date", patentdir)
						sub2files, err := ioutil.ReadDir(dirPath + "/" + patentType + "/" + patentdir)
						if err != nil {
							log.Fatal(err)
						}
						for _, f := range sub2files {
							patentzip := f.Name()

							if strings.Contains(patentzip, ".zip") || strings.Contains(patentzip, ".ZIP") {
								//log.Println("file-zip", patentzip)
								src := dirPath + "/" + patentType + "/" + patentdir + "/" + patentzip

								if strings.Contains(patentType, "IMGS-30-S") {
									// 解压 压缩包至 output 目录
									outputS := output + "/30-S/" + patentdir + "/"
									//log.Println("解压中...", patentdir, endTime[0])
									if patentdir <= endTime[0] {
										log.Println("该压缩包已经处理过了，跳过。")
									} else {
										err = Unzip(src, outputS)
										if err != nil {
											log.Println(err, "解压失败手动处理=====", src)
										}
									}

								} else if strings.Contains(patentType, "TXTS-10-A") {
									outputA := output + "/10-A/" + patentdir + "/"
									//log.Println("解压中...")
									if patentdir <= endTime[1] {
										log.Println("该压缩包已经处理过了，跳过。")
									} else {
										err = Unzip(src, outputA)
										if err != nil {
											log.Println(err, "解压失败手动处理=====", src)
										}
									}

								} else if strings.Contains(patentType, "TXTS-10-B") {
									outputB := output + "/10-B/" + patentdir + "/"
									//log.Println("解压中...")
									if patentdir <= endTime[2] {
										log.Println("该压缩包已经处理过了，跳过。")
									} else {
										err = Unzip(src, outputB)
										if err != nil {
											log.Println(err, "解压失败手动处理=====", src)
										}
									}

								} else if strings.Contains(patentType, "TXTS-20-U") {
									outputU := output + "/20-U/" + patentdir + "/"
									//log.Println("解压中...")
									if patentdir <= endTime[2] {
										log.Println("该压缩包已经处理过了，跳过。")
									} else {
										err = Unzip(src, outputU)
										if err != nil {
											log.Println(err, "解压失败手动处理=====", src)
										}
									}

								}

							}

						}
					}
				}
			}
		}

	}
	log.Println("xml file has been extracted!")

	return err

}

func findXML(output string) error {
	outputArr := []string{"/30-S", "/10-A", "/10-B", "/20-U"}
	for i, v := range outputArr {
		eTime := ""
		output := output + v
		log.Println(output)
		err := HandleWalk(output, i)
		if err != nil {
			return err
		}
		// 删除解压文件
		eTime, err = removeDIR(output)
		if err != nil {
			return err
		}

		endTime[i] = eTime
	}
	log.Println("解析结束，下次从这里开始", endTime)

	return nil

}

func removeDIR(output string) (string, error) {
	// 解析完，清理下原始数据 output/30-S/日期目录
	lastDir := ""
	files, err := ioutil.ReadDir(output)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			log.Println("deleting ...", f.Name())
			err := os.RemoveAll(output + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			lastDir = f.Name()
		}
	}
	fmt.Printf("Delete %s \n", output)
	return lastDir, nil
}

func HandleWalk(output string, patentIndex int) error {
	err := filepath.Walk(output, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, "XML") {
			// parse xml
			switch patentIndex {
			case 0:
				err := parse.Par0Xml(path, output, patentIndex, EngineHadoopGORM)
				if err != nil {
					return err
				}
			case 1:
				err := parse.Par1Xml(path, output, patentIndex, EngineHadoopGORM)
				if err != nil {
					return err
				}

			case 2:
				err := parse.Par1Xml(path, output, patentIndex, EngineHadoopGORM)
				if err != nil {
					return err
				}
			case 3:
				err := parse.Par1Xml(path, output, patentIndex, EngineHadoopGORM)
				if err != nil {
					return err
				}
			}

		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Unzip will unzip zip file and return unzip files dir path
func Unzip(zipFilePath string, output string) error {
	// 1. create tempDir to save unzip files
	log.Println("path", zipFilePath, output)
	err := os.MkdirAll(output, 777)
	if err != nil {
		return err
	}
	err = pkg.UnZip(zipFilePath, output)
	if err != nil {
		return err
	}
	return nil
}
