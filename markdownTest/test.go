package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		fmt.Println(file.Name())
		ext := path.Ext(file.Name())
		println(ext)
	}
}
