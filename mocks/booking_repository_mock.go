package mocks

import (
	"p2-ip-hotel-rental/models"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) CountOverlappingBookings(
	suiteID uint,
	checkIn, checkOut time.Time,
) (int64, error) {

	args := m.Called(suiteID, checkIn, checkOut)

	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBookingRepository) Create(booking *models.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingRepository) FindByUserID(userID uint) ([]models.Booking, error) {
	args := m.Called(userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]models.Booking), args.Error(1)
}
