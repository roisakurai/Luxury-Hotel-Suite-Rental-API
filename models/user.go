package models

import "time"

type User struct {
	ID            uint    `gorm:"primaryKey"`
	Email         string  `gorm:"unique;not null"`
	Password      string  `gorm:"not null"`
	DepositAmount float64 `gorm:"default:0"`
	CreatedAt     time.Time
}
