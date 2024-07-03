package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"size:64"`
	Username string `gorm:"size:64;unique;not null"`
	Email    string `gorm:"size:255;unique;index;not null"`
	Password string `gorm:"size:64;not null"`
}
type UserResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"fullName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
}
