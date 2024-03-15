package models

import (
	"time"
)

// User is a struct to represent a user for gorm
type User struct {
	ID                     uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email                  string    `json:"email" gorm:"unique;not null"`
	Password               string    `json:"-" gorm:"not null;"`
	Verified               bool      `json:"verified" gorm:"default:false"`
	VerificationCode       string    `json:"-" gorm:"not null;size:6"`
	VerificationCodeExpiry time.Time `json:"verification_code_expiry" gorm:"not null;"`
	VerificationAt         time.Time `json:"verification_at" gorm:"default:null"`
}
