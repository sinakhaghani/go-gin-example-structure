package config

import (
	"fmt"
	"go-gin-example-structure/src/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ ", GetMessage("DotenvLoadError")+":", err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ ", GetMessage("DatabaseConnectionError")+":", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ ", GetMessage("DatabaseMigrationError")+":", err)
	}

	log.Println("✅", GetMessage("DatabaseConnectionSuccess"))
}
