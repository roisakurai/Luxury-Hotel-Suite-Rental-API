package service

import (
	"testing"
	"time"

	"p2-ip-hotel-rental/mocks"
	"p2-ip-hotel-rental/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBooking_Success(t *testing.T) {

	userRepo := new(mocks.MockUserRepository)
	suiteRepo := new(mocks.MockSuiteRepository)
	bookingRepo := new(mocks.MockBookingRepository)
	txRepo := new(mocks.MockTransactionRepository)
	emailMock := new(mocks.MockEmailService)

	service := NewBookingService(nil, userRepo, suiteRepo, bookingRepo, txRepo, emailMock)

	emailMock.On("SendEmail", mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	userRepo.On("FindByID", uint(1)).
		Return(&models.User{ID: 1, DepositAmount: 10000000}, nil)

	suiteRepo.On("FindByID", uint(1)).
		Return(&models.Suite{
			ID:            1,
			PricePerNight: 1000000,
			Stock:         5,
		}, nil)

	bookingRepo.On("CountOverlappingBookings", mock.Anything, mock.Anything, mock.Anything).
		Return(int64(0), nil)

	bookingRepo.On("Create", mock.Anything).Return(nil)

	booking, err := service.CreateBooking(
		1,
		1,
		time.Now(),
		time.Now().Add(24*time.Hour),
	)

	assert.Nil(t, err)
	assert.NotNil(t, booking)

	bookingRepo.AssertExpectations(t)
}

func TestGetBookingReport(t *testing.T) {
	bookingRepo := new(mocks.MockBookingRepository)

	service := NewBookingService(nil, nil, nil, bookingRepo, nil, nil)

	mockData := []models.Booking{
		{ID: 1, UserID: 1},
	}

	bookingRepo.On("FindByUserID", uint(1)).
		Return(mockData, nil)

	result, err := service.GetBookingReport(1)

	assert.Nil(t, err)
	assert.Len(t, result, 1)
}

func TestTopUp_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockTxRepo := new(mocks.MockTransactionRepository)

	emailMock := new(mocks.MockEmailService)

	service := NewUserService(mockUserRepo, mockTxRepo, emailMock)

	mockUserRepo.On("UpdateDeposit", uint(1), 100000.0).Return(nil)
	mockTxRepo.On("Create", mock.Anything).Return(nil)

	err := service.TopUp(1, 100000)

	assert.Nil(t, err)
}
