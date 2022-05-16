package domain

type Todo struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string `json:"name" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
}
