package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/danierj/training/to-do-app/api/models"
	"github.com/jinzhu/gorm"
)

func BodyParser(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	return body, err
}

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func ClearTable(db *gorm.DB) {
	db.DropTableIfExists(&models.Todo{})
	db.AutoMigrate(&models.Todo{})
}

func AddTodos(db *gorm.DB, count int) []models.Todo {
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
