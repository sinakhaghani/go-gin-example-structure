package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example-structure/app/validators"

	"go-gin-example-structure/app/common/utils/i18n"
	"go-gin-example-structure/app/repositories"
	"net/http"
)

type UserController struct {
	userRepo repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func Example(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "example message",
		"data":    "",
	})
}

/*
get user from database
*/
func (r *UserController) GetUsers(c *gin.Context) {

	users, err := r.userRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}

	message := i18n.Translate(c, "UsersList", "User list") // from translation src/common/utils/i18n

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    users,
	})
}

/*
create user in database
*/
func (r *UserController) CreateUser(c *gin.Context) {

	var input validators.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.userRepo.CreateUser(input)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
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
