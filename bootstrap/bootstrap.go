package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	customValidator "go-gin-example-structure/app/common/validators"

	"go-gin-example-structure/config"
	"go-gin-example-structure/routes"
)

func SetupApp() *gin.Engine {
	config.InitI18n()
	config.InitDatabase()

	// رجیستر ولیدیتورها
	customValidator.RegisterDBForValidators(config.DB)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("exists", customValidator.ExistsIn)
		if err != nil {
			return nil
		}
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	return r
}
