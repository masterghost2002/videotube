package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	CreateAt time.Time
	UpdateAt time.Time
}
