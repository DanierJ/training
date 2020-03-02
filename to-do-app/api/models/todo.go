package models

import (
	"github.com/jinzhu/gorm"
)

// Todo describes how a todo task should look like
type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(30)" json:"title"`
	Description string `json:"description"`
}

// NewTodo creates a new todo
func NewTodo(todo Todo) error {
	db := Connect()
	defer db.Close()
	rs := db.Create(&todo)
	return rs.Error
}

// FindAll finds all todos
func FindAll() (interface{}, error) {
	db := Connect()
	defer db.Close()
	rs := db.Find(&[]Todo{})
	return rs.Value, rs.Error
}
