package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/danierj/training/to-do-app/api/models"
	"github.com/danierj/training/to-do-app/utils"
)

// GetTodos controller to find all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := models.FindAll()

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, todos, http.StatusOK)

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

	err = models.NewTodo(todo)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, todo, http.StatusCreated)
}
