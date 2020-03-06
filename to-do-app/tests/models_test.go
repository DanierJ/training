package tests

import (
	"log"
	"testing"

	"github.com/danierj/training/to-do-app/api/models"
)

var todo models.Todo

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
	clearTable()

	todo = models.Todo{Title: "", Description: ""}

	rs, err := todo.NewTodo()

	if err == nil {
		t.Errorf("Expected to have an error like 'You can't have empty values' and rs empty. Instead got rs(%v) and error (%v)", rs, err)
	}
}

func TestNewTodoWithValidValues(t *testing.T) {
	clearTable()

	todo = models.Todo{Title: "Valid title", Description: "Valid description"}

	rs, err := todo.NewTodo()

	if err != nil {
		t.Errorf("Expected to have a new todo created. Got rs(%v) want (%v)", rs, todo)
	}
}

func TestNewTodoWithTooLongValues(t *testing.T) {
	clearTable()

	todo = models.Todo{Title: "Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long title", Description: "Tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo long descrption"}

	_, err := todo.NewTodo()

	if errMsg := err.Error(); err != nil && errMsg != "Title is too long. Please write 30 char or less" {
		t.Errorf("Expected to have an error msg like 'You have too long values' and rs empty. Got (%v) want 'Title is too long. Please write 30 char or less'", err)
	}
}

func TestFindAll(t *testing.T) {
	clearTable()
	addTodos(5)
	rs, err := todo.FindAll()
	if err != nil {
		t.Errorf("Expected to have a slice of Todos. Got (%v) want (%v)", err, rs)
	}
}

func TestFindExistingTodoById(t *testing.T) {
	clearTable()
	td := addTodos(1)

	todo, err := td[0].FindByID(1)

	if todo.ID != td[0].ID {
		t.Errorf("Expected todo with ID (%v) to be found. Got (%v) want (%v)", td[0].ID, todo.ID, td[0].ID)
	}

	if err != nil {
		t.Errorf("Expected to fetch todo with id #1. Got (%v).", todo)
	}

}

func TestFindNonExistingTodoById(t *testing.T) {
	clearTable()
	td := addTodos(1)

	todo, err := td[0].FindByID(100)

	if todo.ID == td[0].ID {
		t.Errorf("Expected no todo to be found. Got (%v) want (%v)", todo.ID, err)
	}

	if err == nil {
		t.Errorf("Expected error message like 'Todo not found'. Got (%v) and error msg (%v)", todo, err)
	}

}

func TestUpdateExistingTodoWithValidValue(t *testing.T) {
	clearTable()
	td := addTodos(6) // Find todo

	todo := models.Todo{ID: 5, Title: "New Title", Description: "New Description"}

	updatedTodo, err := todo.UpdateTodo(uint64(todo.ID))

	if err != nil {
		log.Fatal(err)
	}

	if updatedTodo.Title == td[0].Title {
		t.Errorf("Expected todo title to be different. Got (%v) want (%v)", updatedTodo.Title, todo.Title)
	}

	if updatedTodo.Description == td[0].Description {
		t.Errorf("Expected todo description to be different. Got (%v) want (%v)", updatedTodo.Description, todo.Description)
	}

}

func TestUpdateExistingTodoWithEmptyValue(t *testing.T) {
	clearTable()
	addTodos(1) // Find todo

	todo := models.Todo{ID: 1, Title: "", Description: ""}

	updatedTodo, err := todo.UpdateTodo(uint64(todo.ID))

	if err == nil {
		t.Errorf("Expected to have an error like 'You can't have empty values' and rs empty. Instead got rs(%v) and error (%v)", updatedTodo, err)
	}

}

func TestDeleteExistingTodo(t *testing.T) {
	clearTable()
	td := addTodos(6)

	ra, err := todo.DeleteTodo(td[0].ID)
	rowsAffected, ok := ra.(int64)

	if err != nil {
		t.Errorf("Expected the operation to be successful. Got (%v)", err)
	}

	todo, err := td[0].FindByID(uint64(td[0].ID))

	if ok && rowsAffected == 0 && err == nil {
		t.Errorf("Expected to have todo with ID [1] deleted. Got todo with ID (%v)", todo.ID)
	}

	if ok && rowsAffected > 1 && err == nil {
		t.Errorf("Expected to have just one todo (todo with ID [1]) deleted. Got (%v) numbers of rows affected", rowsAffected)
	}

}

func TestDeleteNonExistingTodo(t *testing.T) {
	clearTable()

	ra, err := todo.DeleteTodo(1)

	rowsAffected, ok := ra.(int64)

	if err != nil {
		t.Errorf("Expected the operation to be successful. Got (%v)", err)
	}

	if ok && rowsAffected > 0 {
		t.Errorf("Expected to have no todo deleted. Got (%v) rows affected", rowsAffected)
	}

}
