package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-gin-example-structure/config"
	"go-gin-example-structure/database/models"
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
	localizerAny, exists := c.Get("localizer")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "localizer not set"})
		return
	}
	localizer := localizerAny.(*i18n.Localizer)

	//get users
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	//message
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "UsersList",
	})
	if err != nil {
		msg = "Users list" // fallback
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    users,
	})
}

func CreateUser(c *gin.Context) {

	localizerAny, exists := c.Get("localizer")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "localizer not set"})
		return
	}
	localizer := localizerAny.(*i18n.Localizer)

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

	//message
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CreateUserSuccess",
	})
	if err != nil || msg == "" {
		msg = "User created" // fallback message
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": msg,
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
