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
	"github.com/fatih/color"
	"time"
)

func main() {
	start := time.Now()

	fileName := flag.String("file", "", "The name of your .txt file containing data to be inserted into the template.")
	dirName := flag.String("dir", ".", "The name of your .txt file containing data to be inserted into the template.")
	flag.Parse()

	var files []string
	ext := ".txt"
	if *fileName != "" {
		files = append(files, *fileName)
	} else {
		files = findFilesRec(*dirName, ext)
	}

	var fileSizes float64
	for _, fileName := range files {
		data := readFile(fileName+ext)
		fileSizes += createAndSaveFile(fileName, data)
	}
	s := ""
	if len(files) > 1 {
		s = "s"
	}
	boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	exeTime := time.Since(start).Seconds()
	fmt.Printf("%s Generated %s page%s (%.1fkB total) in %.2f seconds.\n", boldGreen("Success!"), bold(len(files)), s, fileSizes, exeTime)
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
	return files
}

type siteData struct {
    Content string
	Title string
}

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

func createAndSaveFile(filename string, data siteData) float64 {
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
	return float64(len(bytesToWrite)) / float64(1000)
}

func exit(msg string) {
    fmt.Println(msg)
    os.Exit(1)
}
