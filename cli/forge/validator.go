package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func MakeValidator(path string) {
	dir := "src/validators"
	filePath := fmt.Sprintf("%s/%s.go", dir, path)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("validator already exists.")
		return
	}

	fileName := filepath.Base(path)

	tmpl := `package validators


type ExampleInput struct {}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("validator").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(fileName),
	})

	fmt.Println("validator created:", filePath)
}
