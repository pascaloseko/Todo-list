package controllers
import (
	"to-do-list/models"
)

var todos = []models.Todo{
	{
		Title: "My first todo",
		Completed: false,
		UserID: 1,
	},
	{
		Title: "Another todo",
		Completed: true,
		UserID: 2,
	},
}

func setup(){
	TodoDeleteAll()
}