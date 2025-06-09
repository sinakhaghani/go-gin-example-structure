package repositories

import (
	"gorm.io/gorm"

	"go-gin-example-structure/app/models"
	"go-gin-example-structure/app/validators"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	CreateUser(input validators.CreateUserInput) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(input validators.CreateUserInput) (models.User, error) {

	// Password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = r.db.Create(&user).Error

	return user, err
}
