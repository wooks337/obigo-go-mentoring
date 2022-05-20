package domain

import "gorm.io/gorm"

type MajorDepartment struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Major struct {
	gorm.Model
	Name              string `gorm:"not null"`
	MajorDepartmentId int
	MajorDepartment   MajorDepartment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students          []Student
	Professor         []Professor
}

type Student struct {
	gorm.Model
	Name               string `gorm:"not null"`
	MajorId            int
	Major              Major `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age                int
	ClassRegistrations []ClassRegistration
}

type Professor struct {
	gorm.Model
	Name    string `gorm:"not null"`
	MajorId int
	Major   Major `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Classes []Class
}

type Class struct {
	gorm.Model
	Name               string `gorm:"not null"`
	Credit             int    `gorm:"not null"`
	ProfessorId        int
	Professor          Professor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClassRegistrations []ClassRegistration
}

type ClassRegistration struct {
	gorm.Model
	StudentId int
	Student   Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClassId   int
	Class     Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Score struct {
	gorm.Model
	MiddleScore         int
	FinalScore          int
	TotalScore          float64
	ClassRegistrationId int
	ClassRegistration   ClassRegistration `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
