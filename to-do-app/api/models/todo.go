package models

import (
	"errors"
)

// Todo describes how a todo task should look like
type Todo struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Title       string `gorm:"type:varchar(30);not null" json:"title"`
	Description string `json:"description"`
	Error       string `gorm:"-" json:"error"`
}

// NewTodo creates a new todo
func NewTodo(todo Todo) (interface{}, error) {
	db := Connect()
	defer db.Close()
	rs := db.Create(&todo)

	if rs.Error != nil {
		err := rs.Error
		return Todo{Error: err.Error()}, rs.Error
	}

	return rs.Value, nil
}

// FindAll finds all todos
func FindAll() (interface{}, error) {
	db := Connect()
	defer db.Close()
	rs := db.Find(&[]Todo{})

	if rs.Error != nil {
		err := rs.Error
		return Todo{Error: err.Error()}, rs.Error
	}

	return rs.Value, rs.Error
}

// FindByID finds one record by id
func FindByID(id int) (interface{}, error) {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", id).First(&Todo{})

	if rs.Error != nil {
		err := rs.Error
		return Todo{Error: err.Error()}, rs.Error
	}

	return rs.Value, rs.Error

}

// BeforeSave it's a hook
func (t *Todo) BeforeSave() (err error) {

	switch {
	case t.IsEmpty():
		err = errors.New("Can't save empty data")

	case t.TooLong():
		err = errors.New("Title is too long. Please write 30 char or less")
	}
	return
}

// IsEmpty checks wether a todo field is empty or not
func (t Todo) IsEmpty() bool {

	if t.Title == "" || t.Description == "" {
		return true
	}

	return false
}

// TooLong checks wether a todo field is too long
func (t Todo) TooLong() bool {

	if len(t.Title) > 30 {
		return true
	}

	return false
}
