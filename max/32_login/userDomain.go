package loginMod

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null", json:"username"`
	Password string `gorm:"not null", json:"password"`
	Name     string `gorm:"not null", json:"name"`
	Age      int    `gorm:"not null", json:"age"`
	Email    string `gorm:"not null", json:"email"`
}

type SignupUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Email    string `json:"email"`
}
