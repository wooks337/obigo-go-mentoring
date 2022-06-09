package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"unique;not null" json:"userID"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
}
