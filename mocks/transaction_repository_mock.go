package mocks

import (
	"p2-ip-hotel-rental/models"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(tx *models.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}
