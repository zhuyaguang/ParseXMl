package main

import (
	"archive/zip"
	"flag"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io"
	"io/ioutil"
	"log"
	"os"
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
	extracting1Xml(*dataAdd, *outputAdd)

	err :=findXML(*outputAdd)
	if err!=nil{
		fmt.Println(err.Error())
	}


	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

}

func extracting1Xml(dirPath string, output string) error {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// 办理第一层
		if strings.Contains(f.Name(), "TXTS") || strings.Contains(f.Name(), "IMGS-30-S") {
			patentType := f.Name()
			fmt.Println("file-type", patentType)
			if f.IsDir() {
				subfiles, err := ioutil.ReadDir(dirPath + "/" + patentType)
				if err != nil {
					log.Fatal(err)
				}
				for _, f := range subfiles {
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
								if strings.Contains(patentType, "IMGS-30-S") {
									Unzip(src, output+"/30-S/"+patentdir)
								} else if strings.Contains(patentType, "TXTS-10-A") {
									Unzip(src, output+"/10-A/"+patentdir)
								} else if strings.Contains(patentType, "TXTS-10-B") {
									Unzip(src, output+"/10-B/"+patentdir)
								} else if strings.Contains(patentType, "TXTS-20-U") {
									Unzip(src, output+"/20-U/"+patentdir)
								}

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

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())

		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func findXML(output string) error {
	outputArr := []string{"/30-S", "/10-A", "/10-B", "/20-U"}

	for i, v := range outputArr {
		output := output + v
		fmt.Println(output)
		err := HandleWalk(output, i)
		if err != nil {
			return err
		}
	}

	return nil

}
func HandleWalk(output string, patentIndex int) error {
	err := filepath.Walk(output, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, "XML") {
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

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func convertXML2Json(xmlPath, output, filename string) {
	// convert xml to json
	xmlFile, err := os.Open(xmlPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened xml file")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml := strings.NewReader(string(byteValue))
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	f, err := os.Create(output + "/" + filename)
	defer f.Close()

	n3, err := f.WriteString(json.String())
	fmt.Printf("wrote %d bytes\n", n3)
}
