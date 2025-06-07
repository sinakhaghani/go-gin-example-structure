package controllers

import (
	"github.com/gin-gonic/gin"

	"go-gin-example-structure/config"
	"go-gin-example-structure/src/common/utils/i18n"
	"go-gin-example-structure/src/models"
	"go-gin-example-structure/src/validators"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Example(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "example message",
		"data":    "",
	})
}

/*
get user from database
*/
func GetUsers(c *gin.Context) {

	//get users
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	message := i18n.Translate(c, "UsersList", "User list") // from translation src/common/utils/i18n

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    users,
	})
}

func CreateUser(c *gin.Context) {

	// src/validators/user (validation form)
	var input validators.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in password hashing"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// Save to database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	message := i18n.Translate(c, "CreateUserSuccess", "User created")

	c.JSON(http.StatusCreated, gin.H{
		"message": message,
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
