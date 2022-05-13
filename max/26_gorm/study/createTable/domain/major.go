package domain

import "gorm.io/gorm"

type Major struct {
	gorm.Model
	Name              string          `gorm:"not null"`
	MajorDepartmentId int             `gorm:"not null"`
	MajorDepartment   MajorDepartment `gorm:"constraint:OnUpdate:CASCADE;"`
	Students          []Student
}
