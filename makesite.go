package main

import (
	"io/ioutil"
	"text/template"
	"bytes"
    "flag"
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"time"
	"github.com/fatih/color"
	"github.com/gomarkdown/markdown"
)

func main() {
	start := time.Now()
	fileName := flag.String("file", "",
		"The name of a file containing data to be inserted into the template, without the extension!")
	dirName := flag.String("dir", ".",
		"The directory you'd like to search for data files in.")
	extName := flag.String("ext", ".txt",
		"Extension for data file(s) you'd like to use.")
	tmplName := flag.String("tmpl", "template",
		"A .tmpl file to insert data into.")
	flag.Parse()

	if (*extName)[0] != '.' {
		*extName = "." + (*extName)
	}
	if strings.LastIndex(*tmplName, ".") < 0 {
		*tmplName += ".tmpl"
	}

	var files []string
	if *fileName != "" {
		files = append(files, *fileName)
	} else {
		files = findFilesRec(*dirName, *extName)
	}

	var fileSizes float64
	for _, fileName := range files {
		data := readFile(fileName, *extName)
		fileSizes += createAndSaveFile(*tmplName, fileName, data)
	}
	s := ""
	if len(files) > 1 {
		s = "s"
	}
	boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	exeTime := time.Since(start).Seconds()
	fmt.Printf("%s Generated %s page%s (%.1fkB total) in %.2f seconds.\n",
		boldGreen("Success!"), bold(len(files)), s, fileSizes, exeTime)
}

type siteData struct {
    Content string
	Title string
}

func findFilesRec(dir string, ext string) []string {
	var files []string

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
		    if err != nil {
		        return err
		    }
			if filepath.Ext(path) == ext {
			    files = append(files, path[:len(path)-len(ext)])
			}
		    return nil
		})
	checkErr(err, fmt.Sprintf("Crawling for %s files in %s!\n", ext, dir))
	return files
}

func readFile(filename string, ext string) siteData {
	fileBytes, err := ioutil.ReadFile(filename + ext)
	checkErr(err, fmt.Sprintf("Error while reading %s%s!\n", filename, ext))
	title := filename
	var content []byte
	if string(fileBytes[:6]) == "TITLE:" {
		for i := 0; i < len(fileBytes); i++ {
	        if fileBytes[i] == '\n' {
				title = string(fileBytes[6:i])
				content = []byte(string(fileBytes[i+1:]))
				break
			}
	    }
	} else {
		i := strings.LastIndex(filename,"/")
		if i >= 0 {
			title = string(filename[i+1:])
		}
		content = fileBytes
	}

	return siteData{Content: string(markdown.ToHTML(content, nil, nil)),
		Title: strings.TrimSpace(title)}
}

func createAndSaveFile(tmplName string, filename string, data siteData) float64 {
	paths := []string{tmplName}
	buffer := new(bytes.Buffer)
	t := template.Must(template.New(tmplName).ParseFiles(paths...))
	err := t.Execute(buffer, data)
	checkErr(err, fmt.Sprintf("Error while parsing template %s!\n", tmplName))
	bytesToWrite := []byte(buffer.String())
	err = ioutil.WriteFile(filename + ".html", bytesToWrite, 0644)
	checkErr(err, fmt.Sprintf("Error while inserting data into %s.html!\n", filename))
	return float64(len(bytesToWrite)) / float64(1000)
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	    panic(err)
	}
}
