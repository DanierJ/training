package models

import (
	"errors"
)

// Todo describes how a todo task should look like
type Todo struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	Title       string `gorm:"type:varchar(30);not null" json:"title"`
	Description string `json:"description"`
	Error       string `gorm:"-" json:"error"`
}

// NewTodo creates a new todo
func (t *Todo) NewTodo() (*Todo, error) {
	db := Connect()
	defer db.Close()
	rs := db.Create(&t)

	if rs.Error != nil {
		err := rs.Error
		return &Todo{Error: err.Error()}, rs.Error
	}

	return t, nil
}

// FindAll finds all todos
func (t *Todo) FindAll() (*[]Todo, error) {
	db := Connect()
	defer db.Close()

	todos := []Todo{}
	rs := db.Find(&todos)

	if rs.Error != nil {
		return &[]Todo{}, rs.Error
	}

	return &todos, rs.Error
}

// FindByID finds one record by id
func (t *Todo) FindByID(id uint64) (*Todo, error) {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", id).First(&t)

	if rs.Error != nil {
		err := rs.Error
		return &Todo{Error: err.Error()}, rs.Error
	}

	return t, rs.Error

}

// UpdateTodo updates record if found
func (t *Todo) UpdateTodo(id uint64) (*Todo, error) {
	db := Connect()
	defer db.Close()
	rs := db.Model(&t).Where("id = ?", id).Updates(
		map[string]interface{}{
			"title":       t.Title,
			"description": t.Description,
		},
	)
	if rs.Error != nil {
		err := rs.Error
		return &Todo{Error: err.Error()}, rs.Error
	}

	return t, rs.Error
}

func (t *Todo) DeleteTodo(id uint64) (interface{}, error) {
	db := Connect()
	defer db.Close()

	rs := db.Where("id = ?", id).Delete(&t)

	if err := rs.Error; err != nil {
		return &Todo{Error: err.Error()}, err
	}

	return rs.RowsAffected, nil

}

// BeforeSave it's a gorm hook that gets executed before saving a record
func (t *Todo) BeforeSave() (err error) {
	return validateInputs(t)
}

// BeforeUpdate it's a gorm hook that gets executed before updating a record
func (t *Todo) BeforeUpdate() (err error) {
	return validateInputs(t)
}

// IsEmpty checks wether a todo field is empty or not
func (t Todo) isEmpty() bool {

	if t.Title == "" || t.Description == "" {
		return true
	}

	return false
}

// TooLong checks wether a todo field is too long
func (t Todo) tooLong() bool {

	if len(t.Title) > 30 {
		return true
	}

	return false
}

func validateInputs(t *Todo) (err error) {
	switch {
	case t.isEmpty():
		err = errors.New("Can't save empty data")

	case t.tooLong():
		err = errors.New("Title is too long. Please write 30 char or less")
	}
	return
}
