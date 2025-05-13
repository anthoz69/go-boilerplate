package utils

import (
	"strings"

	"github.com/anthoz69/salepage-api/shared/constants"
	"github.com/gofiber/fiber/v2"
)

func GetLang(ctx *fiber.Ctx) string {
	lang := strings.ToUpper(ctx.Get("Accept-Language"))
	if lang == "" || (lang != constants.LangEN && lang != constants.LangTH) {
		return constants.LangTH
	}
	return lang
}
