package Service

import (
	"gorm.io/gorm"
	"obigo-go-mentoring/jamie/03GORM/GormRestAPI/database"
)

//전체 리스트 조회
func GetTodoList(db *gorm.DB) (todo []database.Todo, err error) {
	res := db.Order("ID").Find(&todo)
	err = res.Error

	return todo, err
}

func PostTodoList(db *gorm.DB) (todo []database.Todo, err error) {

}
