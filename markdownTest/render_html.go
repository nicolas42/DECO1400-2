package main

import (
	"io/ioutil"
	"path"

	"gitlab.com/golang-commonmark/markdown"
)

func main() {
	infile := "Balloon-Technology-58567049-fcad-4a80-ac26-829f073554b2.md"
	ext := path.Ext(infile)
	outfile := infile[0:len(infile)-len(ext)] + ".html"

	md := markdown.New(markdown.XHTMLOutput(true))
	bytes, _ := ioutil.ReadFile(infile)
	ioutil.WriteFile(outfile, []byte(md.RenderToString(bytes)), 0644)
}

//	fmt.Println(md.RenderToString([]byte("Header\n===\nText")))
