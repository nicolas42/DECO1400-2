package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

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
			ioutil.WriteFile(outfile, []byte(md.RenderToString(bytes)), 0644)

		}
	}
}

//	fmt.Println(md.RenderToString([]byte("Header\n===\nText")))
