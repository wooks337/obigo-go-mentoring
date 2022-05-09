package domain

type Actor struct {
	Actor_id    int
	First_name  string
	Last_name   string
	Last_update string
}

type Film_actor struct {
	Actor_id    int
	Film_id     int
	Last_update string
}

type Film struct {
	Film_id              int
	Title                string
	Description          string
	Release_year         string
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

//뒤에 s 붙는 문제 해결
func (Actor) TableName() string {
	return "actor"
}
func (Film) TableName() string {
	return "film"
}
