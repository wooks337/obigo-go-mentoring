package domain

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string `gorm:"not null"`
	StudentId *int   `gorm:"unique"`
	//StudentId int
	Age     int
	MajorId int   `gorm:"not null"`
	Major   Major `gorm:"foreignKey:MajorId;constraint:OnUpdate:CASCADE;"`
}
