package user

import (
	domain "github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/internal/core/ports"
	"github.com/anthoz69/salepage-api/internal/middleware"
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
		return utils.NewErrorResponse(c, constants.ErrCodeInvalidInput, err.Error())
	}
	if err := h.useCase.CreateUser(&user); err != nil {
		return utils.NewErrorResponse(c, constants.ErrCodeBusinessLogic, err.Error())
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return utils.NewErrorResponse(c, constants.ErrCodeUnexpectedError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"token":   token,
	})
}
