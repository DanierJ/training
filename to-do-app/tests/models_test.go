package tests

import (
	"testing"

	"github.com/danierj/training/to-do-app/api/models"
)

func TestDBConnection(t *testing.T) {
	if models.Connect() == nil {
		t.Errorf("Expetected to stablish db connection. Instead got nil")
	}
}

func TestAutoMigrations(t *testing.T) {
	if !models.Connect().HasTable("todos") {
		t.Errorf("Expected to have todos table created")
	}
}

func TestNewTodoWithEmptyValues(t *testing.T) {
	rs, err := models.NewTodo(models.Todo{Title: "", Description: ""})

	if err == nil {
		t.Errorf("Expected to have an error like 'You can't have empty values' and rs empty. Instead got rs(%v) and error (%v)", rs, err)
	}
}

func TestNewTodoWithValidValues(t *testing.T) {
	rs, err := models.NewTodo(models.Todo{Title: "Valid title", Description: "Valid description"})

	if err != nil {
		t.Errorf("Expected to have a new todo created. Got rs(%v) want (%v)", rs, models.Todo{Title: "Valid title", Description: "Valid description"})
	}
}

func TestNewTodoWithTooLongValues(t *testing.T) {
	_, err := models.NewTodo(models.Todo{Title: "Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long title", Description: "Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long descrption"})

	if errMsg := err.Error(); err != nil && errMsg != "Title is too long. Please write 30 char or less" {
		t.Errorf("Expected to have an error msg like 'You have too long values' and rs empty. Got (%v) want 'Title is too long. Please write 30 char or less'", err)
	}
}

func TestFindAll(t *testing.T) {
	rs, err := models.FindAll()
	if err != nil {
		t.Errorf("Expected to have a slice of Todos. Got (%v) want (%v)", err, rs)
	}
}

func TestFindExistingTodoById(t *testing.T) {
	todo, err := models.FindByID(1)

	if err != nil {
		t.Errorf("Expected to fetch todo with id #1. Got (%v).", todo)
	}

}

func TestFindNonExistingTodoById(t *testing.T) {
	todo, err := models.FindByID(100)

	if err == nil {
		t.Errorf("Expected error message like 'Todo not found'. Got (%v) and error msg (%v)", todo, err)
	}

}
