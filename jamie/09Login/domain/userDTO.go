package domain

import "gorm.io/gorm"

type User struct {
	gorm.DB
	UserID   string
	Password string
	Name     string
	Email    string
}
