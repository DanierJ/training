package routes

import (
	"github.com/danierj/training/to-do-app/api/actions"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", actions.GetTodos).Methods("GET")
	r.HandleFunc("/todos", actions.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", actions.GetTodo).Methods("GET")
	r.HandleFunc("/todos", actions.PostTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", actions.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", actions.DeleteTodo).Methods("DELETE")

	return r
}
