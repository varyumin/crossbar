package path

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ListFiles(path string) []string {
	var fileClean []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".gz" || filepath.Ext(file.Name()) == ".tar" {
			fileClean = append(fileClean, filepath.Clean(fmt.Sprintf("%s/%s", path, file.Name())))
		}
	}
	return fileClean
}
