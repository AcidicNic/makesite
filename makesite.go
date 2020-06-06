package main

import (
	"io/ioutil"
	"html/template"
	"bytes"
    "flag"
)


func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
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

type siteData struct {
    Content string
}

func main() {
	fileName := flag.String("file", "content", "The name of your .txt file containing data to be inserted into the template.")
	flag.Parse()

	data := siteData{Content: readFile(*fileName+".txt")}

	createAndSaveFile(*fileName, data)
}
