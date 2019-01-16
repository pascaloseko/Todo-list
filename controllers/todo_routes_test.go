package controllers
import (
	"to-do-list/models"
)

import "testing"

func TodoDeleteAll() (err error){
	db := models.GetDB()
	defer db.Close()
	statement := "delete from todos"
	_, err = db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func TestCreateTodo(t *testing.T){

}