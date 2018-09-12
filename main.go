package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const codeTemplate = `
  package main
  import "log"

  func main() {
    log.Print("%s")
  }
`

const fileName = "code.go"

func main() {
	path, err := getFilePath()
	log.Printf("path: %s", path)
	if err != nil {
		panic(err)
	}
	log.Print(path)
	createFile(path, "hello!!!")
	cmd := exec.Command("go", "run", path)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Print(err.Error())
	}
}

func getFilePath() (string, error) {
	dirPath, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	return filepath.Join(dirPath, fileName), nil
}

func createFile(fileName string, str string) error {
	code := []byte(fmt.Sprintf(codeTemplate, str))
	err := ioutil.WriteFile(fileName, code, 0644)
	if err != nil {
		return err
	}
	return nil
}
