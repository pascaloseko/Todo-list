package main

import (
	"net/http"
	"time"
	"to-do-list/app"
	"to-do-list/controllers"
	u "to-do-list/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// attach JWT middleware
	router.Use(app.JWTAuthentication)
	//prints out the configuration details
	u.P("Todo List", u.Version(), "started at ", u.Config.Address)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/create/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/me/todo", controllers.GetTodo).Methods("GET")

	server := &http.Server{
		Addr:           u.Config.Address,
		Handler:        router,
		ReadTimeout:    time.Duration(u.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(u.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}