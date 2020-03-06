package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danierj/training/to-do-app/api/models"
	"github.com/danierj/training/to-do-app/api/utils"
	"github.com/gorilla/mux"
)

// GetTodos controller to find all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	Todo := models.Todo{}

	todos, err := Todo.FindAll()
	status := setStatusCode(http.StatusOK, err)
	utils.ToJson(w, todos, status)

}

// PostTodo controller to create a new Todo
func PostTodo(w http.ResponseWriter, r *http.Request) {
	saveTodo(w, r, -1, http.StatusCreated)
}

// GetTodo controller to get one Todo
func GetTodo(w http.ResponseWriter, r *http.Request) {

	todo := models.Todo{}
	id := getID(w, r)

	if id == 0 {
		return
	}

	t, err := todo.FindByID(id)

	status := setStatusCode(http.StatusOK, err)

	utils.ToJson(w, t, status)

}

// UpdateTodo controller to update Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	id := getID(w, r)

	if id == 0 {
		return
	}

	saveTodo(w, r, int(id), http.StatusOK)

}

// DeleteTodo controller to delete Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	id := getID(w, r)

	if id == 0 {
		return
	}

	_, err := todo.DeleteTodo(id)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, nil, http.StatusNoContent)

}

func setStatusCode(status int, err error) int {

	if err != nil {
		switch err.Error() {
		case "Can't save empty data", "Title is too long. Please write 30 char or less":
			return http.StatusUnprocessableEntity
		case "record not found":
			return http.StatusNotFound
		}
	}
	return status
}

func saveTodo(w http.ResponseWriter, r *http.Request, id int, status int) {
	body, err := utils.BodyParser(r)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	var savedTodo *models.Todo

	if id > 0 { // UPDATE
		savedTodo, err = todo.UpdateTodo(uint64(id))
	}

	savedTodo, err = todo.NewTodo()

	status = setStatusCode(status, err)
	utils.ToJson(w, savedTodo, status)

}

func getID(w http.ResponseWriter, r *http.Request) uint64 {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return 0
	}

	return id
}
