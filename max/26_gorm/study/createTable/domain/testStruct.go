package domain

type TestCredit struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	TestUserId int
}

type TestUser struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	TestCredit TestCredit
}
