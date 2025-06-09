package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func MakeValidator(path string) {
	dir := "app/validations"
	filePath := fmt.Sprintf("%s/%s.go", dir, path)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("validation already exists.")
		return
	}

	fileName := filepath.Base(path)

	tmpl := `package validations


type ExampleInput struct {}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("validation").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(fileName),
	})

	fmt.Println("validation created:", filePath)
}
