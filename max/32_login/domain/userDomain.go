package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null" json:"username"`
	Password     string `gorm:"not null" json:"password"`
	Name         string `gorm:"not null" json:"name"`
	Age          int    `gorm:"not null" json:"age"`
	Email        string `gorm:"not null" json:"email"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InfoUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

type ClaimUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
