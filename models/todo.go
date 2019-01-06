package models

import (
	"fmt"
	u "to-do-list/utils"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (todo *Todo) Validate() (map[string]interface{}, bool) {
	if todo.Title == "" {
		return u.Message(false, "Todo tile should be on the payload"), false
	}

	if todo.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

// Create Todo
func (todo *Todo) Create() map[string]interface{} {
	if resp, ok := todo.Validate(); !ok {
		return resp
	}

	GetDB().Create(todo)
	resp := u.Message(true, "success")
	resp["todo"] = todo
	return resp
}

// GetTodo ...
func GetTodo(id uint) *Todo {
	todo := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil
	}
	return todo
}

// GetTodos ...
func GetTodos(user uint) []*Todo {
	todos := make([]*Todo, 0)
	err := GetDB().Table("todos").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return todos
}
