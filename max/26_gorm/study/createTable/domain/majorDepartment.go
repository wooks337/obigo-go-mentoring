package domain

import "gorm.io/gorm"

type MajorDepartment struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Majors []Major
}
