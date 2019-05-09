package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gitlab.com/golang-commonmark/markdown"
)

func main() {

	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		ext := path.Ext(file.Name())
		if ext == ".md" {
			fmt.Println(file.Name())

			infile := file.Name()
			ext := path.Ext(infile)
			outfile := infile[0:len(infile)-len(ext)] + ".html"

			md := markdown.New(markdown.XHTMLOutput(true))
			bytes, _ := ioutil.ReadFile(infile)
			html := md.RenderToString(bytes)
			html = strings.ReplaceAll(html, ".md", ".html")
			html = header + html
			html = html + footer
			ioutil.WriteFile(outfile, []byte(html), 0644)

		}
	}
}

var header = `<!doctype html>
<html>
	<head>
		<link rel="stylesheet" type="text/css" href="style.css">
	</head>
	<body>
`

var footer = `</body>
</html>
`
