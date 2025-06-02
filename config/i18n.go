package config

import (
	"encoding/json"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

func InitI18n() {
	Bundle = i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := Bundle.LoadMessageFile("translations/en.json")
	if err != nil {
		log.Fatal("❌ Error loading en.json:", err)
	}

	_, err = Bundle.LoadMessageFile("translations/fa.json")
	if err != nil {
		log.Fatal("❌ Error loading fa.json:", err)
	}
}

func GetMessage(id string) string {
	// localizer := i18n.NewLocalizer(Bundle, "fa")
	localizer := i18n.NewLocalizer(Bundle, "en")
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: id})
	if err != nil {
		return id
	}
	return msg
}
