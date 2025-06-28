package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
)

// uniqueIn: custom validator tag
func UniqueIn(fl validator.FieldLevel) bool {
	log.Println("✅ UniqueIn validator called")

	// tag format: "table.column"
	param := fl.Param() // e.g. "reservations.reservation_code"
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

	return err == nil && count == 0
}
