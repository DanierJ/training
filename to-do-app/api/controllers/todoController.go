package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/danierj/training/to-do-app/api/models"
	"github.com/danierj/training/to-do-app/utils"
	"github.com/gorilla/mux"
)

// GetTodos controller to find all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.FindAll()
	status := setStatusCode(http.StatusOK, err)
	utils.ToJson(w, todos, status)

}

// PostTodo controller to create a new Todo
func PostTodo(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var todo models.Todo
	err := json.Unmarshal(body, &todo)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	t, err := models.NewTodo(todo)
	status := setStatusCode(http.StatusCreated, err)

	utils.ToJson(w, t, status)
}

// GetTodo controller to get one Todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	t, err := models.FindByID(id)

	status := setStatusCode(http.StatusOK, err)

	utils.ToJson(w, t, status)

}

// UpdateTodo controller to update Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Not implemented")
	log.Fatal(err)

}

// DeleteTodo controller to delete Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Not implemented")
	log.Fatal(err)

}

func setStatusCode(status int, err error) int {

	if err != nil {
		switch err.Error() {
		case "Can't save empty data", "Title is too long. Please write 30 char or less":
			return http.StatusBadRequest
		case "record not found":
			return http.StatusNotFound
		}
	}
	return status
}
