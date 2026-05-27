package repository

import (
	"p2-ip-hotel-rental/models"
	"time"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking *models.Booking) error
	CountOverlappingBookings(suiteID uint, checkIn, checkOut time.Time) (int64, error)
	FindByUserID(userID uint) ([]models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CountOverlappingBookings(
	suiteID uint,
	checkIn, checkOut time.Time,
) (int64, error) {

	var count int64

	err := r.db.Model(&models.Booking{}).
		Where("suite_id = ?", suiteID).
		Where("status = ?", "booked").
		Where("check_in_date < ? AND check_out_date > ?", checkOut, checkIn).
		Count(&count).Error

	return count, err
}

func (r *bookingRepository) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) FindByUserID(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}
