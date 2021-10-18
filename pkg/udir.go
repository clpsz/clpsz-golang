package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type DirEntityType string

var DIR DirEntityType = "dir"
var FILE DirEntityType = "file"
var ALL DirEntityType = "all"

func ListFileInDir(dir string, fullPath bool) []string {
	return ListEntityInDir(dir, FILE, fullPath)
}

func ListSubDirInDir(dir string, fullPath bool) []string {
	return ListEntityInDir(dir, DIR, fullPath)
}

func ListEntityInDir(root string, _type DirEntityType, fullPath bool) []string {
	res := make([]string, 0)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if _type == DIR && !f.IsDir() {
			continue
		}
		if _type == FILE && f.IsDir() {
			continue
		}

		if fullPath {
			res = append(res, path.Join(root, f.Name()))
		} else {
			res = append(res, f.Name())
		}
	}
	return res
}

func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0744)
		if err != nil {
			return err
		}
	}

	return nil
}
