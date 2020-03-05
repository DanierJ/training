package routes

import (
	"github.com/danierj/training/to-do-app/api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos", controllers.PostTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	return r
}
