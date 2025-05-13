package user

import (
	domain "github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/internal/core/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}
