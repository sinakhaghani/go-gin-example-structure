package forge

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func MakeMiddleware(path string) {
	dir := "app/middlewares"
	filePath := fmt.Sprintf("%s/%s.go", dir, path)
	fullDir := filepath.Dir(filePath)

	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		os.MkdirAll(fullDir, os.ModePerm)
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Middlewares already exists.")
		return
	}

	fileName := filepath.Base(path)

	tmpl := `package middlewares

	import (
		"github.com/gin-gonic/gin"
	)
	
	func {{.FuncName}}() gin.HandlerFunc {
		return func(c *gin.Context) {
			// Before request
	
			c.Next()
	
			// After request
		}
	}`

	f, _ := os.Create(filePath)
	defer f.Close()

	t := template.Must(template.New("middleware").Parse(tmpl))
	t.Execute(f, map[string]string{
		"FuncName": ToCamelCase(fileName),
	})

	fmt.Println("middleware created:", filePath)
}
