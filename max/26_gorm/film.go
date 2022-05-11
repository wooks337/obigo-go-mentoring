package _6_gorm

type Film struct {
	Film_id              int `gorm:"primaryKey;autoIncrement:true"`
	Title                string
	Description          string
	Release_year         int
	Language_id          int
	Original_language_id int
	Rental_duration      int
	Rental_rate          float64
	Length               int
	Replacement_cost     float64
	Rating               string
	Special_features     string
	Last_update          string
}

func (Film) TableName() string {
	return "film"
}
