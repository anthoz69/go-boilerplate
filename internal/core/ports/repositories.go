package ports

import "github.com/anthoz69/salepage-api/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
}
