# 🔗 Static Site Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/acidicnic/makesite)](https://goreportcard.com/report/github.com/acidicnic/makesite)

### Examples

[README.md Demo](https://acidicnic.github.io/makesite/) created from this readme!

[Animal Crossing Demo](https://acidicnic.github.io/makesite/test_io/animalcrossing)
 - datafile: [animalcrossing.txt](https://github.com/AcidicNic/makesite/blob/master/test_io/animalcrossing.txt)

[Cheese Demo](https://acidicnic.github.io/makesite/test_io/cheese)
 - datafile: [cheese.md](https://github.com/AcidicNic/makesite/blob/master/test_io/cheese.md)


## Getting Started

```bash
git clone https://github.com/AcidicNic/makesite.git
cd makesite

go run makesite.go
```
#### What will this do? (All Default Flags)

 - Crawl the current directory and all subdirectories for .txt files
 - Create HTML files for each .txt file using template.tmpl
 - New HTML files will have the same filename as the .txt data file, they will be stored in the same directory as their source .txt file.

## Flags

```bash
// Will make .html files for each .txt file in the folder /test_io using my_tmpl.tmpl
go run makesite.go -dir="/test_io" -tmpl="my_tmpl"

// Will only make an .html file for README.md in the folder /test_io
go run makesite.go -file="readme" -ext="md"
```

- __dir__ _string_
 - The directory you'd like to search for data files in. (_default "."_)
- __ext__ _string_
	- Extension for data file(s) you'd like to use. (_default ".txt"_)
- __file__ _string_
	- The name of a file containing data to be inserted into the template, without the extension!
    - __Only use this flag if you want to create a single html file.__
- __tmpl__ _string_
	- A .tmpl file to insert data into. (_default "template"_)

## Data File Format

- If first line of the file starts with "TITLE:", the text on that line will be the title in the head and navbar of the HTML file outputted. Otherwise, the filename will be the title.
- You can use a mix of markdown, plain text, and HTML (if you want to...) in your data file. This will all be inserted into the body of the template.
