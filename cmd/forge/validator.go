package forge

import (
	"fmt"
	"os"
	"text/template"
)

func MakeValidator(name string) {
	dir := "src/validators"
	filePath := fmt.Sprintf("%s/%s.go", dir, name)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("validator already exists.")
		return
	}

	tmpl := `package validators


type ExampleInput struct {}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("validator").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(name),
	})

	fmt.Println("validator created:", filePath)
}
