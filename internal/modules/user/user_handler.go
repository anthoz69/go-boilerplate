package user

import (
	domain "github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/internal/core/ports"
	"github.com/anthoz69/salepage-api/shared/constants"
	"github.com/anthoz69/salepage-api/shared/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase ports.UserUseCase
}

func NewUserHandler(useCase ports.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return utils.NewErrorResponse(c, constants.ErrCode0001, err.Error())
	}
	if err := h.useCase.CreateUser(&user); err != nil {
		return utils.NewErrorResponse(c, constants.ErrCode0001, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
