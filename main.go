package main

import (
	"go-gin-example-structure/bootstrap"
)

func main() {
	r := bootstrap.SetupRouter()
	r.Run(":8080")
}
