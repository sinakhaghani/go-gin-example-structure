package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-example-structure/src/common/container"
	"go-gin-example-structure/src/controllers"
	"go-gin-example-structure/src/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	di := container.NewContainer() //Dependency Injection Container

	api := r.Group("/api", middlewares.LocaleMiddleware())
	api.GET("/users", di.UserController.GetUsers)
	api.POST("/users", di.UserController.CreateUser)

	//Routes that require authentication(JWT)
	protected := api.Group("/", middlewares.AuthMiddleware())
	protected.GET("/example", controllers.Example)

}
