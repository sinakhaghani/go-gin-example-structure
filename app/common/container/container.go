package container

import (
	userController "go-gin-example-structure/app/controllers"
	"go-gin-example-structure/app/repositories"
	"go-gin-example-structure/config"
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
