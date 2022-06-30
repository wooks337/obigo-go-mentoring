package domain

import "gorm.io/gorm"

type TestInfo struct {
	gorm.Model
	Name   string
	Min    int    `gorm:"default:0"`
	Sec    int    `gorm:"default:0"`
	Status string `gorm:"default:'on'"`
}

func (TestInfo) TableName() string {
	return "test_info"
}
