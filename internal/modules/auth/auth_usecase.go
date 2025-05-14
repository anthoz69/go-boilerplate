package auth

import (
	"github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/internal/core/ports"
)

type authUseCase struct {
	userRepo ports.UserRepository
}

func NewAuthUseCase(userRepo ports.UserRepository) ports.AuthUseCase {
	return &authUseCase{userRepo: userRepo}
}

func (u *authUseCase) GetMe() (*domain.User, error) {
	return u.userRepo.GetUserById("1")
}
