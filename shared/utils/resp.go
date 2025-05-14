package utils

import (
	"github.com/anthoz69/salepage-api/shared/constants"
	"github.com/gofiber/fiber/v2"
)

func NewSuccessResponse(ctx *fiber.Ctx, message string) error {
	lang := GetLang(ctx)
	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"code":      constants.SuccessCode,
			"codeLabel": constants.MessageMap[constants.SuccessCode].CodeLabel[lang],
			"message":   message,
		})
}
