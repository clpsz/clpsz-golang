package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func ListFileInDir(dir string) []string {
	res := make([]string, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		res = append(res, path.Join(dir, f.Name()))
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
