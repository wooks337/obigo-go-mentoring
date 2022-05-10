package domain

type Student struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	StudentId int
	MajorId   int `gorm:"not null"`
	Age       int
}

func (Student) TableName() string {
	return "student"
}

type MajorDepartment struct {
	Id   int    `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"not null"`
}

func (MajorDepartment) TableName() string {
	return "major_department"
}
