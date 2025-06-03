package forge

import (
	"fmt"
	"os"
	"text/template"
)

func MakeModel(name string) {
	dir := "database/models"
	filePath := fmt.Sprintf("%s/%s.go", dir, name)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
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
