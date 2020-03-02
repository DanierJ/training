package api

import (
	"github.com/danierj/training/to-do-app/api/models"
)

func Run() {
	models.AutoMigrations()
	initTodos()
	initServer(8080)

}

func initTodos() {
	models.NewTodo(models.Todo{Title: "Client meeting", Description: "Meeting with a client"})
	models.NewTodo(models.Todo{Title: "Design app interface", Description: "Desining app interface"})
	models.NewTodo(models.Todo{Title: "Learn Node.js", Description: "Learning to code Node.js"})
	models.NewTodo(models.Todo{Title: "Bring the puppy out XD", Description: "Puppy needs to go out"})
}
