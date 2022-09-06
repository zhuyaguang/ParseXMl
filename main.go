package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"patentExtr/pkg"
	"patentExtr/pkg/parse"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	dataAdd := flag.String("data", "/data/sipo", "source data address")
	outputAdd := flag.String("output", "/data/output", "output xml address")
	flag.Parse()

	fmt.Println(*dataAdd, *outputAdd)

	start := time.Now()
	// Code to measure
	duration := time.Since(start)

	// 把专利数据解压到 output 目录
	// extractingXml(*dataAdd, *outputAdd)

	fmt.Println(duration)

	// 解析 XML 数据成 json 文件
	err := findXML(*outputAdd)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

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
			fmt.Println("file-type", patentType)
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
						fmt.Println("file-date", patentdir)
						sub2files, err := ioutil.ReadDir(dirPath + "/" + patentType + "/" + patentdir)
						if err != nil {
							log.Fatal(err)
						}
						for _, f := range sub2files {
							patentzip := f.Name()

							if strings.Contains(patentzip, ".zip") || strings.Contains(patentzip, ".ZIP") {
								fmt.Println("file-zip", patentzip)
								src := dirPath + "/" + patentType + "/" + patentdir + "/" + patentzip
								//if strings.Contains(patentType, "IMGS-30-S") {
								//	// 解压 压缩包至 output 目录
								//	outputS := output + "/30-S/" + patentdir + "/"
								//	err = Unzip(src, outputS)
								//	if err != nil {
								//		log.Fatal(err)
								//	}
								//	//err := HandleWalk(output, 0)
								//	//if err != nil {
								//	//	return err
								//	//}
								//
								//} //else
								if strings.Contains(patentType, "TXTS-10-A") {
									outputA := output + "/10-A/" + patentdir + "/"
									fmt.Println("解压中...")
									err = Unzip(src, outputA)
									if err != nil {
										fmt.Println(err, "解压失败手动处理=====", src)
									}
									//err := HandleWalk(output, 1)
									//if err != nil {
									//	return err
									//}
								}
								//else if strings.Contains(patentType, "TXTS-10-B") {
								//	outputB := output + "/10-B/" + patentdir + "/"
								//	err = Unzip(src, outputB)
								//	if err != nil {
								//		log.Fatal(err)
								//	}
								//	//err := HandleWalk(output, 2)
								//	//if err != nil {
								//	//	return err
								//	//}
								//} else if strings.Contains(patentType, "TXTS-20-U") {
								//	outputU := output + "/20-U/" + patentdir + "/"
								//	err = Unzip(src, outputU)
								//	if err != nil {
								//		log.Fatal(err)
								//	}
								//	//err := HandleWalk(output, 3)
								//	//if err != nil {
								//	//	return err
								//	//}
								//}

							}

						}
					}
				}
			}
		}

	}
	fmt.Println("xml file has been extracted!")

	return err

}

func findXML(output string) error {
	outputArr := []string{"/30-S", "/10-A", "/10-B", "/20-U"}

	for i, v := range outputArr {
		if v == "/10-A" {
			output := output + v
			fmt.Println(output)
			err := HandleWalk(output, i)
			if err != nil {
				return err
			}
			removeDIR(output)
		}

	}

	return nil

}

func removeDIR(output string) error {
	// 解析完，清理下原始数据 output/30-S/日期目录
	files, err := ioutil.ReadDir(output)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			err := os.RemoveAll(output + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Printf("Delete %s \n", output)
	return err
}

func HandleWalk(output string, patentIndex int) error {
	err := filepath.Walk(output, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path,"XML"){
			fmt.Printf("%s \n", path)
			// parse xml
			switch patentIndex {
			case 0:
				err := parse.Par0Xml(path, output, patentIndex)
				if err != nil {
					return err
				}
			case 1:
				err := parse.Par1Xml(path, output, patentIndex)
				if err != nil {
					return err
				}

			case 2:
				err := parse.Par1Xml(path, output, patentIndex)
				if err != nil {
					return err
				}
			case 3:
				err := parse.Par1Xml(path, output, patentIndex)
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
	fmt.Println("path", zipFilePath, output)
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
