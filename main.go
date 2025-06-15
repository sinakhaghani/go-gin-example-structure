package main

import (
	"go-gin-example-structure/bootstrap"
)

func main() {
	r := bootstrap.SetupApp()
	r.Run(":8080")
}
