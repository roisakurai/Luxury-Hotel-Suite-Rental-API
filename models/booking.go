package models

import "time"

type Booking struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	SuiteID      uint
	CheckInDate  time.Time
	CheckOutDate time.Time
	TotalPrice   float64
	Status       string
	CreatedAt    time.Time
}
