package pkg

import (
"bufio"
	"io/ioutil"
	"log"
"os"
)

func ReadFileAsList(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func WriteContentToFile(content string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

func ReadFileContent(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text
}
