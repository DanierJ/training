package routes

import (
	"github.com/danierj/training/to-do-app/api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.PostTodo).Methods("POST")

	return r
}
