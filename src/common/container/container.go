package container

import (
	"go-gin-example-structure/config"
	userController "go-gin-example-structure/src/controllers"
	"go-gin-example-structure/src/repositories"
)

type Container struct {
	UserController *userController.UserController
	// در اینجا بقیه کنترلرها رو هم می‌تونی اضافه کنی
}

func NewContainer() *Container {
	db := config.DB

	// Repositories
	userRepo := repositories.NewUserRepository(db)

	// Controllers
	userCtrl := userController.NewUserController(userRepo)

	return &Container{
		UserController: userCtrl,
	}
}
