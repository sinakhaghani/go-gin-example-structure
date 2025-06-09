package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"strings"
)

var DB *gorm.DB

// RegisterDBForValidators ثبت اتصال دیتابیس برای استفاده در اعتبارسنجی
func RegisterDBForValidators(db *gorm.DB) {
	DB = db
}

// existsIn: custom validator tag
func ExistsIn(fl validator.FieldLevel) bool {
	log.Println("✅ ExistsIn validator called")

	// tag format: "table,column"
	param := fl.Param() // e.g. "cities,id"
	parts := strings.Split(param, ".")
	if len(parts) != 2 {
		return false
	}

	table := parts[0]
	column := parts[1]
	value := fl.Field().Interface()

	var count int64
	query := fmt.Sprintf("%s = ?", column)
	err := DB.Table(table).Where(query, value).Count(&count).Error

	return err == nil && count > 0
}
