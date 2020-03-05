package tests

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
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var r *mux.Router = routes.NewRouter()
var db *gorm.DB = models.Connect()

func TestGetTodos(t *testing.T) {
	clearTable()

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
	clearTable()

	payload := []byte(`{"title": "test todo", "description": "this is a test description"}`)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "test todo" {
		t.Errorf("Expected todo title to be 'test todo'. Got '%v'", m["title"])
	}
	if m["description"] != "this is a test description" {
		t.Errorf("Expected todo title to be 'this is a test description'. Got '%v'", m["description"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected todo ID to be '1'. Got '%v'", m["id"])
	}

}

func TestEmptyTable(t *testing.T) {
	clearTable()

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
	clearTable()
	var todos = addTodos(1)

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var fetchedTodo map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &fetchedTodo)

	if fetchedTodo["id"] != 1.0 {
		t.Errorf("Expected todo #1 to be fetched. Got %v - Want %v", fetchedTodo, todos[0])
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

func clearTable() {
	db.DropTableIfExists(&models.Todo{})
	db.AutoMigrate(&models.Todo{})
}

func addTodos(count int) []models.Todo {
	var todos []models.Todo

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		todos = append(todos, models.Todo{Title: "Todo #" + strconv.Itoa(i+1), Description: "This is the todo #" + strconv.Itoa(i+1)})

		db.Create(&todos[i])
	}

	return todos
}
