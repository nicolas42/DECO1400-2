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
	os.Chdir("content")
	dir, _ := os.Getwd()
	render_html(dir)
}

func render_html(dir string) {

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		ext := path.Ext(file.Name())
		if ext == ".md" {
			fmt.Println(file.Name())

			// Prepare Files
			infile := file.Name()
			ext := path.Ext(infile)
			outfile := infile[0:len(infile)-len(ext)] + ".html"

			// Generate HTML
			md := markdown.New(markdown.XHTMLOutput(true))
			bytes, _ := ioutil.ReadFile(infile)
			html := md.RenderToString(bytes)

			// Fenaigle
			html = strings.ReplaceAll(html, ".md", ".html")
			html = header + html
			html = html + footer

			// Output File
			ioutil.WriteFile(outfile, []byte(html), 0644)

		}
	}
}

var header = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" type="text/css" href="style.css">
		<script type="text/javascript" src="tools.js"></script>
        <script src="jquery-1.4.2.js"></script>
        <script src="jquery.hotkeys.js"></script>
		
	</head>
	<body>
<div class="inner">


		<h1 class="title">
			<a href="javascript:showHide('nav.container');">&#9776; &nbsp;</a> 
			Getting Around the Solar System
		</h1>
		<nav class="container">
			<div><a href="Where-to-go.html">Where To Go</a></div>			
			<div><a href="The-Balloon-and-the-Elephant.html">The Elephant and the Balloon</a></div>
			<div><a href="Balloon-Technology.html">Balloon Technology</a></div>
			<div><a href="Elephants.html">Elephants</a></div>
			<div><a href="Explosions.html">Explosions</a></div>
			<div><a href="Getting-Around-the-Solar-System.html">Getting Around the Solar System</a></div>
			<div><a href="Gravity.html">Gravity</a></div>
			<div><a href="On-Dreams.html">On Dreams</a></div>
			<div><a href="Orbit.html">Orbit</a></div>
			<div><a href="Space.html">Space</a></div>
		</nav>

`

var footer = `
<nav class="nav-footer">
		<a id="prevPage" href="Balloon-Technology.html"><div class="chevron left"></div></a>
		<a id="nextPage" href="The-Balloon-and-the-Elephant.html"><div class="chevron right"></div></a>
	</nav>


</div><!-- inner -->
<footer>
	<a href="Where-to-go.html">Where To Go</a>			
	<a href="The-Balloon-and-the-Elephant.html">The Elephant and the Balloon</a>
	<a href="Balloon-Technology.html">Balloon Technology</a>
	<a href="Elephants.html">Elephants</a>
	<a href="Explosions.html">Explosions</a>
	<a href="Getting-Around-the-Solar-System.html">Getting Around the Solar System</a>
	<a href="Gravity.html">Gravity</a>
	<a href="On-Dreams.html">On Dreams</a>
	<a href="Orbit.html">Orbit</a>
	<a href="Space.html">Space</a>
</footer>

<script>
	function nextPage() {
		document.getElementById('nextPage').click();
	}
	function prevPage() {
		document.getElementById('prevPage').click();
	}
	$(document).bind('keyup', 'left', prevPage);
	$(document).bind('keyup', 'right', nextPage);
</script>
</body>
</html>
`
