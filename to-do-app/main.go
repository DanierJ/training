package main

import (
	"github.com/danierj/training/to-do-app/todo"
)

func main() {
	initServer(getTodos())
}

func getTodos() []todo.Todo {
	return []todo.Todo{
		todo.New("Client meeting", "Meeting with a client"),
		todo.New("Design app interface", "Desining app interface"),
		todo.New("Learn Node.js", "Learning to code Node.js"),
		todo.New("Bring the puppy out XD", "Puppy needs to go out"),
	}
}
