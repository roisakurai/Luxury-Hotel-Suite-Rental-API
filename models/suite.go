package models

import "time"

type Suite struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Category      string
	PricePerNight float64
	Stock         int
	CreatedAt     time.Time
}
