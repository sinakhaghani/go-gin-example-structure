package i18n

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetLocalizer(c *gin.Context) (*i18n.Localizer, bool) {
	localizerAny, exists := c.Get("localizer")
	if !exists {
		return nil, false
	}
	localizer, ok := localizerAny.(*i18n.Localizer)
	return localizer, ok
}

func Translate(c *gin.Context, messageID string, defaultMsg string) string {
	localizer, ok := GetLocalizer(c)
	if !ok {
		return defaultMsg
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil || msg == "" {
		return defaultMsg
	}
	return msg
}
