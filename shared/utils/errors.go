package utils

import (
	"fmt"

	"github.com/anthoz69/salepage-api/shared/constants"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code      string `json:"code"`
	CodeLabel string `json:"codeLabel"`
	Message   string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.CodeLabel, e.Message)
}

func NewErrorResponse(ctx *fiber.Ctx, code string, message string) error {
	lang := GetLang(ctx)
	errCode, ok := constants.MessageMap[code]
	if !ok {
		errCode = constants.MessageMap[constants.ErrCodeUnexpectedError]
		return ctx.Status(errCode.HTTPStatus).
			JSON(fiber.Map{
				"code":      constants.ErrCodeUnexpectedError,
				"codeLabel": errCode.CodeLabel[lang],
				"message":   message,
			})
	}
	return ctx.Status(errCode.HTTPStatus).
		JSON(fiber.Map{
			"code":      errCode.Code,
			"codeLabel": errCode.CodeLabel[lang],
			"message":   message,
		})
}
