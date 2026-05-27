package models

import (
	"time"
)

type UserResponse struct {
	ID            uint    `json:"id"`
	Email         string  `json:"email"`
	DepositAmount float64 `json:"deposit_amount"`
}

type BookingResponse struct {
	ID           uint      `json:"id"`
	SuiteID      uint      `json:"suite_id"`
	CheckInDate  time.Time `json:"check_in"`
	CheckOutDate time.Time `json:"check_out"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`
}

func ToUserResponse(user *User) UserResponse {
	return UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
	}
}

func ToBookingResponse(b *Booking) BookingResponse {
	return BookingResponse{
		ID:           b.ID,
		SuiteID:      b.SuiteID,
		CheckInDate:  b.CheckInDate,
		CheckOutDate: b.CheckOutDate,
		TotalPrice:   b.TotalPrice,
		Status:       b.Status,
	}
}

type SuiteResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	PricePerNight float64 `json:"price_per_night"`
	Stock         int     `json:"stock"`
}

func ToSuiteResponse(s Suite) SuiteResponse {
	return SuiteResponse{
		ID:            s.ID,
		Name:          s.Name,
		Category:      s.Category,
		PricePerNight: s.PricePerNight,
		Stock:         s.Stock,
	}
}

type UserProfileResponse struct {
	ID            uint    `json:"id"`
	Email         string  `json:"email"`
	DepositAmount float64 `json:"deposit_amount"`
}

func ToUserProfileResponse(u *User) UserProfileResponse {
	return UserProfileResponse{
		ID:            u.ID,
		Email:         u.Email,
		DepositAmount: u.DepositAmount,
	}
}

type SuitesResponse struct {
	Message string          `json:"message"`
	Data    []SuiteResponse `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Token   string       `json:"token"`
}

type UserProfileWrapper struct {
	Message string              `json:"message"`
	Data    UserProfileResponse `json:"data"`
}
