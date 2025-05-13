package ports

import "github.com/anthoz69/salepage-api/internal/core/domain"

type UserUseCase interface {
	CreateUser(user *domain.User) error
}
