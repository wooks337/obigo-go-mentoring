package domain

type Dept struct {
	DeptID    uint `gorm:"primaryKey;autoIncrement"`
	DeptName  string
	DeptBuild string
}
type Student struct {
	StuID   uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	Age     int
	Gender  string
	Country string `gorm:"default:'south korea'"`
	DeptID  uint   `gorm:"foreignkey:DeptID"`
}
type Prof struct {
	ProfID  uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	Age     int
	Gender  string
	Country string `gorm:"default:'south korea'"`
	DeptID  uint   `gorm:"foreignkey:DeptID"`
}
