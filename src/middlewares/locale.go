package middlewares

import (
	"go-gin-example-structure/config"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func LocaleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}

		localizer := i18n.NewLocalizer(config.Bundle, lang)
		c.Set("localizer", localizer)
		c.Next()
	}
}
