package database

type Student struct {
	StudentID uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null"`
	Subject   string `json:"subject"`
	Score     int    `json:"score"`
	//SubjectID   uint
	//SubStudents []SubStudent
}

//
//type Prof struct {
//	ProfID    uint `gorm:"primaryKey;autoIncrement"`
//	Name      string
//	Age       int
//	SubjectID uint
//	Subject   Subject
//}
//type SubStudent struct {
//	SubjectID uint `gorm:"primaryKey"`
//	StudentID uint `gorm:"primaryKey"`
//}
//
//type Subject struct {
//	SubjectID   uint `gorm:"primaryKey;autoIncrement"`
//	SubName     string
//	StudentID   uint
//	SubStudents []SubStudent
//}

//func (Subject) TableName() string {
//	return "subject"
//}
//func (Student) TableName() string {
//	return "student"
//}
//func (Prof) TableName() string {
//	return "prof"
//}
