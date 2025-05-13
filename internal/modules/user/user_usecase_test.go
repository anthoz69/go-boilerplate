package user

import (
	"testing"

	"github.com/anthoz69/salepage-api/internal/core/domain"
	"github.com/anthoz69/salepage-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := &domain.User{
		Username: "test",
		Password: "test",
	}

	mockRepo := new(mocks.UserRepository)
	mockRepo.On("CreateUser", user).Return(nil)

	result := NewUserUseCase(mockRepo).CreateUser(user)

	assert.NoError(t, result)
}
