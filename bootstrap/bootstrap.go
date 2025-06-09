package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	customValidator "hotel_guest/src/common/validators"

	"hotel_guest/config"
	"hotel_guest/routes"
)

func SetupRouter() *gin.Engine {
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
