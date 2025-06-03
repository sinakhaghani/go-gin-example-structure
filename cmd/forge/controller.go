package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func MakeController(name string) {
	dir := "src/controllers"
	filePath := fmt.Sprintf("%s/%s.go", dir, name)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Controller already exists.")
		return
	}

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
		"FuncName": ToCamelCase(name),
	})

	fmt.Println("Controller created:", filePath)
}
