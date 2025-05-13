package user

import (
	"github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/internal/core/ports"
)

type userUseCase struct {
	repo ports.UserRepository
}

func NewUserUseCase(repo ports.UserRepository) ports.UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}
