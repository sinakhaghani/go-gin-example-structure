package forge

import (
	"fmt"
	"os"
	"text/template"
)

func MakeController(name string) {
	dir := "src/controllers"
	filePath := fmt.Sprintf("%s/%s.go", dir, name)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
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
