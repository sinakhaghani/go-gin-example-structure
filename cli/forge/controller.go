package forge

import (
	"fmt"
	"os"
	"path/filepath"

	"text/template"
)

func MakeController(path string) {
	dir := "app/controllers"
	filePath := fmt.Sprintf("%s/%s.go", dir, path)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Controller already exists.")
		return
	}

	fileName := filepath.Base(path)

	tmpl := `package controllers

	import (
		"github.com/gin-gonic/gin"
		"net/http"
	)
	
	func {{.FuncName}}(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from {{.FuncName}}",
		})
	}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("controller").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(fileName),
	})

	fmt.Println("Controller created:", filePath)
}
