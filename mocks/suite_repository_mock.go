package mocks

import (
	"p2-ip-hotel-rental/models"

	"github.com/stretchr/testify/mock"
)

type MockSuiteRepository struct {
	mock.Mock
}

func (m *MockSuiteRepository) FindByID(id uint) (*models.Suite, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Suite), args.Error(1)
}

func (m *MockSuiteRepository) FindAll() ([]models.Suite, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Suite), args.Error(1)
}
