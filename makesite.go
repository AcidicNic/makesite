package main

import (
	"io/ioutil"
	"html/template"
	"bytes"
    "flag"
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

func readFile(filename string) siteData {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to %s\n", err))
	}
	var title string
	var content string
	var final int
	for i := 0; i <= 32; i++ {
		final = i
        if fileContents[i] == '\n' {
			break
		}
    }
	title = string(fileContents[0:final])
	content = string(fileContents[final+1:])
	return siteData{Content: strings.TrimSpace(string(content)), Title: strings.TrimSpace(string(title))}
}

func createAndSaveFile(filename string, data siteData) {
	paths := []string{"template.tmpl"}
	buffer := new(bytes.Buffer)
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	err := t.Execute(buffer, data)
	if err != nil {
		panic(err)
	}
	bytesToWrite := []byte(buffer.String())
	err = ioutil.WriteFile(filename + ".html", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func findFilesRec(dir string, ext string) []string {
	var files []string

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
		    if err != nil {
		        return err
		    }
			if filepath.Ext(path) == ext {
			    files = append(files, path[:len(path)-4])
			}
		    return nil
		})
	if err != nil {
	    exit(fmt.Sprintf("Failed to %s\n", err))
	}
	fmt.Println(files)
	return files
}

type siteData struct {
    Content string
	Title string
}

func main() {
	// fileName := flag.String("file", "content", "The name of your .txt file containing data to be inserted into the template.")
	dirName := flag.String("dir", ".", "The name of your .txt file containing data to be inserted into the template.")
	flag.Parse()

	ext := ".txt"

	files := findFilesRec(*dirName, ext)

	for _, fileName := range files {
		data := readFile(fileName+ext)
		createAndSaveFile(fileName, data)
	}
}

func exit(msg string) {
    fmt.Println(msg)
    os.Exit(1)
}
