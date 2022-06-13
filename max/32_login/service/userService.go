package service

import (
	"gorm.io/gorm"
	"loginMod/domain"
)

func Signup(db *gorm.DB, user domain.User) error {

	result := db.Create(&user)

	return result.Error
}

func UsernameDuplicateCheck(db *gorm.DB, username string) bool {

	findUser := domain.User{}
	result := db.Model(&domain.User{}).First(&findUser, "username = ?", username)
	if result.Error != nil {
		return true
	} else {
		return false
	}
}

func FindUserByUsername(db *gorm.DB, username string) (domain.User, error) {
	findUser := domain.User{}
	result := db.Model(&domain.User{}).First(&findUser, "username = ?", username)
	return findUser, result.Error
}
