package api

import (
	"github.com/danierj/training/to-do-app/api/models"
)

func Run() {
	models.AutoMigrations()
	InitTodos()
	InitServer("8080")
}

func InitTodos() {

	var todos []models.Todo
	todos = append(todos,
		models.Todo{Title: "Client meeting", Description: "Meeting with a client"},
		models.Todo{Title: "Design app interface", Description: "Desining app interface"},
		models.Todo{Title: "Learn Node.js", Description: "Learning to code Node.js"},
		models.Todo{Title: "Bring the puppy out XD", Description: "Puppy needs to go out"})

	for _, todo := range todos {
		todo.NewTodo()
	}
}
