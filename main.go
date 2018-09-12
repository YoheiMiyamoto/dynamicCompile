package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const codeTemplate = `
  package main
  import "fmt"

  func main() {
    fmt.Print("%s")
  }
`

const fileName = "code.go"

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	str, err := execute(r.URL.Query().Get("str"))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, str)
}

func execute(str string) (string, error) {
	path, err := getFilePath()
	if err != nil {
		return "", err
	}
	createFile(path, str)
	cmd := exec.Command("go", "run", path)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// err = cmd.Run()
	// if err != nil {
	// 	return err
	// }
	return string(out), nil
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
