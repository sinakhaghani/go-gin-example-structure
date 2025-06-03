package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func MakeModel(name string) {
	dir := "src/models"
	filePath := fmt.Sprintf("%s/%s.go", dir, name)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("model already exists.")
		return
	}

	tmpl := `package models


type ExampleModel struct {}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("model").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(name),
	})

	fmt.Println("model created:", filePath)
}
