package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"unique"`
	Password string // Password is not displayed in JSON output.
}
