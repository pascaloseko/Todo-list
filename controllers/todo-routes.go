package controllers

import (
	"encoding/json"
	"to-do-list/models"
	u "to-do-list/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateTodo ...
var CreateTodo = func (w http.ResponseWriter, r *http.Request){
	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	todo.UserID = user
	resp := todo.Create()
	u.Respond(w, resp)
}

// GetTodo a single todo
var GetTodo = func (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was a problem with your request"))
		return
	}

	data := models.GetTodos(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}