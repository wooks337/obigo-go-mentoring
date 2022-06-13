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

func UpdateRefreshToken(db *gorm.DB, id uint, refreshToken string) bool {

	user := domain.User{}
	res := db.First(&user, id)
	if res.Error != nil {
		return false
	}

	//res = db.Model(&user).Updates(domain.User{RefreshToken: refreshToken})
	res = db.Model(&user).Updates(map[string]interface{}{
		"refresh_token": refreshToken,
	})

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
