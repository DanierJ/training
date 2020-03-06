package actions_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/danierj/training/to-do-app/api"
	"github.com/danierj/training/to-do-app/api/models"
	"github.com/danierj/training/to-do-app/api/routes"
	"github.com/danierj/training/to-do-app/api/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var r *mux.Router = routes.NewRouter()
var db *gorm.DB = models.Connect()

func TestGetTodos(t *testing.T) {
	utils.ClearTable(db)

	api.InitTodos()

	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := 4

	var todos []models.Todo

	json.Unmarshal(response.Body.Bytes(), &todos)

	if len(todos) != expected {
		t.Errorf("Handler returned unexpected number of elements: got %v want %v", len(todos), expected)
	}

}

func TestCreateTodo(t *testing.T) {
	payload := []byte(`{"title": "test todo", "description": "this is a test description"}`)

	expected := models.Todo{ID: 1, Title: "test todo", Description: "this is a test description"}

	saveTodo(t, payload, "POST", "/todos", expected, 1)

}

func TestEmptyTable(t *testing.T) {
	utils.ClearTable(db)

	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var todos []models.Todo

	json.Unmarshal(response.Body.Bytes(), &todos)

	if len(todos) != 0 {
		t.Errorf("Expected an empty array. Got an array with %v elements", len(todos))
	}
}

func TestGetTodo(t *testing.T) {
	utils.ClearTable(db)
	var todos = utils.AddTodos(db, 1)

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var fetchedTodo map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &fetchedTodo)

	if fetchedTodo["id"] != 1.0 {
		t.Errorf("Expected todo #1 to be fetched. Got %v - Want %v", fetchedTodo, todos[0])
	}

}

func TestUpdateTodo(t *testing.T) {

	todoID := 6

	payload := []byte(`{"title": "New title", "description": "New description"}`)

	expected := models.Todo{ID: uint64(todoID), Title: "New title", Description: "New description"}

	saveTodo(t, payload, "PATCH", "/todos/"+strconv.Itoa(6), expected, todoID)

}

func TestDeleteTodo(t *testing.T) {
	utils.ClearTable(db)
	initialTodos := 6
	todoID := 4

	utils.AddTodos(db, initialTodos)

	var foundTodo models.Todo

	req, _ := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(todoID), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNoContent, response.Code)

	rs := db.Where("id = ?", todoID).First(&foundTodo)

	if err := rs.Error; err != nil {
		t.Logf(err.Error())
	}

	if foundTodo.ID == uint64(todoID) {
		t.Errorf("Expected todo with ID (%v) not found", todoID)
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists(tableName string) {
	if !db.HasTable(tableName) {
		log.Fatal("This table doesn't exist")
	}
}

func saveTodo(t *testing.T, payload []byte, method, url string, expected models.Todo, initialTodos int) {
	utils.ClearTable(db)

	var todos []models.Todo
	var expectedStatus = http.StatusCreated

	if method == "PATCH" {
		todos = utils.AddTodos(db, initialTodos)
		expectedStatus = http.StatusOK
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, expectedStatus, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if method == "PATCH" {
		m["id"] = float64(todos[initialTodos-1].ID)
	}

	if m["title"] != expected.Title {
		t.Errorf("Expected todo title to be 'test todo'. Got '%v'", m["title"])
	}

	if m["description"] != expected.Description {
		t.Errorf("Expected todo title to be 'this is a test description'. Got '%v'", m["description"])
	}

	if m["id"] != float64(expected.ID) {
		t.Errorf("Expected todo ID to be (%v). Got (%v)", m["id"], expected.ID)
	}
}
