package service

import (
	"errors"
	"fmt"
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/repository"
	"time"

	"gorm.io/gorm"
)

type BookingService interface {
	CreateBooking(userID, suiteID uint, checkIn, checkOut time.Time) (*models.Booking, error)
	GetBookingReport(userID uint) ([]models.Booking, error)
}

type bookingService struct {
	db           *gorm.DB
	userRepo     repository.UserRepository
	suiteRepo    repository.SuiteRepository
	bookingRepo  repository.BookingRepository
	txRepo       repository.TransactionRepository
	emailService EmailService
}

func NewBookingService(
	db *gorm.DB,
	userRepo repository.UserRepository,
	suiteRepo repository.SuiteRepository,
	bookingRepo repository.BookingRepository,
	txRepo repository.TransactionRepository,
	emailService EmailService,
) BookingService {
	return &bookingService{
		db, userRepo, suiteRepo, bookingRepo, txRepo, emailService,
	}
}

func (s *bookingService) CreateBooking(
	userID, suiteID uint,
	checkIn, checkOut time.Time,
) (*models.Booking, error) {

	if !checkOut.After(checkIn) {
		return nil, errors.New("invalid date range")
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	suite, err := s.suiteRepo.FindByID(suiteID)
	if err != nil {
		return nil, errors.New("suite not found")
	}

	count, err := s.bookingRepo.CountOverlappingBookings(suiteID, checkIn, checkOut)
	if err != nil {
		return nil, err
	}

	if int(count) >= suite.Stock {
		return nil, errors.New("no availability")
	}

	nights := int(checkOut.Sub(checkIn).Hours() / 24)
	totalPrice := float64(nights) * suite.PricePerNight

	if user.DepositAmount < totalPrice {
		return nil, errors.New("insufficient balance")
	}

	var booking *models.Booking

	if s.db == nil {
		b := &models.Booking{
			UserID:       userID,
			SuiteID:      suiteID,
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			TotalPrice:   totalPrice,
			Status:       "booked",
		}

		err := s.bookingRepo.Create(b)
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.User{}).
			Where("id = ?", userID).
			Update("deposit_amount", gorm.Expr("deposit_amount - ?", totalPrice)).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.Transaction{
			UserID: userID,
			Amount: totalPrice,
			Type:   "payment",
		}).Error; err != nil {
			return err
		}

		b := &models.Booking{
			UserID:       userID,
			SuiteID:      suiteID,
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			TotalPrice:   totalPrice,
			Status:       "booked",
		}

		if err := tx.Create(b).Error; err != nil {
			return err
		}

		booking = b
		return nil
	})

	err = s.emailService.SendEmail(
		user.Email,
		"Booking Confirmation",
		fmt.Sprintf("Booking berhasil untuk suite ID %d, total: %.2f", suiteID, totalPrice),
	)

	if err != nil {
		return nil, err
	}

	return booking, nil

}

func (s *bookingService) GetBookingReport(userID uint) ([]models.Booking, error) {
	return s.bookingRepo.FindByUserID(userID)
}
