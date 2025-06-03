package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-example-structure/config"
	"go-gin-example-structure/routes"
	"go-gin-example-structure/src/models"
	"log"
)

func main() {
	r := gin.Default()
	config.InitI18n()
	config.InitDatabase()

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Error in migration: ", err)
	}

	routes.RegisterRoutes(r)
	r.Run(":8080")
}
