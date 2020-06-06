package main

import (
	"io/ioutil"
	"html/template"
	// "os"
	"bytes"
)


func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

type siteData struct {
    Content string
}

func main() {
	data := siteData{Content: readFile("first-post.txt")}

	paths := []string{"template.tmpl"}

	buffer := new(bytes.Buffer)

	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))

	err := t.Execute(buffer, data)
	if err != nil {
		panic(err)
	}

	bytesToWrite := []byte(buffer.String())
	err = ioutil.WriteFile("first-post.html", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}
