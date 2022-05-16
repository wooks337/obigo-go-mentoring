package Service

import (
	"fmt"
	"gorm.io/gorm"
	"todo/domain"
)

func GetTodoList(db *gorm.DB) (todolist []domain.Todo, err error) {

	res := db.Order("ID").Find(&todolist)
	err = res.Error

	return todolist, err
}

func PostTodo(db *gorm.DB, todo domain.Todo) (domain.Todo, error) {
	fmt.Println(todo)
	res := db.Create(&todo)
	return todo, res.Error
}

func RemoveTodo(db *gorm.DB, id int) error {
	res := db.Delete(domain.Todo{}, id)
	if res.RowsAffected == 0 {
		return fmt.Errorf("Not Found")
	}
	return nil
}

func GetTodoById(db *gorm.DB, id int) (todo *domain.Todo) {

	res := db.First(&todo, id)
	if res.Error != nil {
		return nil
	}
	return todo
}

func UpdateTodo(db *gorm.DB, todo domain.Todo) (findTodo domain.Todo, err error) {

	db.Find(&findTodo, todo.ID)
	fmt.Println(todo)
	fmt.Println(findTodo)
	findTodo.Name = todo.Name
	findTodo.Completed = todo.Completed
	res := db.Updates(&findTodo)
	return findTodo, res.Error
}
