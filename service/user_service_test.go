package service

import (
	"errors"
	"testing"

	"p2-ip-hotel-rental/mocks"
	"p2-ip-hotel-rental/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister_EmailExists(t *testing.T) {

	mockRepo := new(mocks.MockUserRepository)
	emailMock := new(mocks.MockEmailService)

	emailMock.On("SendEmail", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := NewUserService(mockRepo, nil, emailMock)

	mockRepo.On("FindByEmail", "test@mail.com").
		Return(&models.User{Email: "test@mail.com"}, nil)

	err := service.Register("test@mail.com", "123456")

	assert.NotNil(t, err)
}

func TestLogin_Success(t *testing.T) {

	mockRepo := new(mocks.MockUserRepository)
	emailMock := new(mocks.MockEmailService)

	service := NewUserService(mockRepo, nil, emailMock)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	user := &models.User{
		ID:       1,
		Email:    "test@mail.com",
		Password: string(hashed),
	}

	mockRepo.On("FindByEmail", "test@mail.com").
		Return(user, nil)

	result, err := service.Login("test@mail.com", "123456")

	assert.Nil(t, err)
	assert.Equal(t, user, result)
}

func TestRegister_Success(t *testing.T) {

	mockRepo := new(mocks.MockUserRepository)
	emailMock := new(mocks.MockEmailService)

	emailMock.On("SendEmail", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := NewUserService(mockRepo, nil, emailMock)

	mockRepo.On("FindByEmail", "test@mail.com").
		Return(nil, errors.New("not found"))

	mockRepo.On("Create", mock.Anything).
		Return(nil)

	err := service.Register("test@mail.com", "123456")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
