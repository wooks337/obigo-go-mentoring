package service

import (
	"gorm.io/gorm"
	"school/domain"
)

func GetMajorList(db *gorm.DB) (majorList []domain.Major) {
	db.Order("ID").Find(&majorList)
	return majorList
}

func GetMajorByPK(db *gorm.DB, id uint) (findMajor *domain.Major) {
	res := db.First(&findMajor, id)
	if res.Error != nil {
		return nil
	}
	return findMajor
}

func AddMajor(db *gorm.DB, major domain.Major) (uint, error) {

	res := db.Create(&major)
	if res.Error != nil {
		return 0, res.Error
	}
	return major.ID, nil
}

func UpdateMajorName(db *gorm.DB, updateMajor domain.Major) error {

	oldMajor := GetMajorByPK(db, updateMajor.ID)
	oldMajor.Name = updateMajor.Name
	res := db.Updates(&oldMajor)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteMajor(db *gorm.DB, id uint) error {
	res := db.Table("major").Delete(id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
