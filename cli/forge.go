package main

import (
	"fmt"
	"go-gin-example-structure/cli/forge"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: forge make:<type> <name>")
		return
	}

	command := os.Args[1]
	name := os.Args[2]

	switch command {
	case "make:controller":
		forge.MakeController(name)
	case "make:middleware":
		forge.MakeMiddleware(name)
	case "make:validator":
		forge.MakeValidator(name)
	case "make:model":
		forge.MakeModel(name)
	case "make:migration":
		forge.MakeMigration(name)

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
