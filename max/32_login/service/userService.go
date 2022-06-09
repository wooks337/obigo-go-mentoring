package service

import (
	"gorm.io/gorm"
	"loginMod"
)

func Signup(db *gorm.DB, user loginMod.User) error {

	result := db.Create(&user)

	return result.Error
}
