package archive

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func printAllMDFiles() {
	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		ext := path.Ext(file.Name())
		if ext == ".md" {
			fmt.Println(file.Name())

		}
	}
}
