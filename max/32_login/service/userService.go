package service

import (
	"gorm.io/gorm"
	"loginMod"
)

func Signup(db *gorm.DB, user loginMod.User) error {

	result := db.Create(&user)

	return result.Error
}

func UsernameDuplicateCheck(db *gorm.DB, username string) bool {

	findUser := loginMod.User{}
	result := db.Model(&loginMod.User{}).First(&findUser, "username = ?", username)
	if result.Error != nil {
		return true
	} else {
		return false
	}
}

func FindUserByUsername(db *gorm.DB, username string) (loginMod.User, error) {
	findUser := loginMod.User{}
	result := db.Model(&loginMod.User{}).First(&findUser, "username = ?", username)
	return findUser, result.Error
}
