package service

import (
	"gorm.io/gorm"
	"school/domain"
)

func GetDepartmentList(db *gorm.DB) (departmentList []domain.MajorDepartment) {

	db.Order("ID").Find(&departmentList)
	return departmentList
}

func GetDepartmentByPK(db *gorm.DB, id uint) (findDepartment *domain.MajorDepartment) {
	res := db.First(&findDepartment, id)
	if res.Error != nil {
		return nil
	}
	return findDepartment
}

func AddDepartment(db *gorm.DB, department domain.MajorDepartment) (uint, error) {

	res := db.Create(&department)
	if res.Error != nil {
		return 0, res.Error
	}
	return department.ID, nil
}

func UpdateDepartmentName(db *gorm.DB, updateDepartment domain.MajorDepartment) error {

	oldDepartment := GetDepartmentByPK(db, updateDepartment.ID)
	oldDepartment.Name = updateDepartment.Name
	res := db.Updates(&oldDepartment)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteDepartment(db *gorm.DB, id uint) error {
	res := db.Table("department").Delete(id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
