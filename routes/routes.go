package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-example-structure/src/controllers"
	"go-gin-example-structure/src/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api", middlewares.LocaleMiddleware())
	api.GET("/users", controllers.GetUsers)
	api.POST("/users", controllers.CreateUser)

	//Routes that require authentication(JWT)
	protected := api.Group("/", middlewares.AuthMiddleware())
	protected.GET("/example", controllers.Example)

}
