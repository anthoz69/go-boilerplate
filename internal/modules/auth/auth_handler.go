package auth

import (
	"github.com/anthoz69/salepage-api/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	useCase ports.AuthUseCase
}

func NewAuthHandler(useCase ports.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase: useCase}
}

func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	user, err := h.useCase.GetMe()
	if err != nil {
		return err
	}
	return c.JSON(user)
}
