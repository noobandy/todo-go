package dao

import (
	"day4/todo/model"
)

type ToDoDAO interface {
	Add(model.ToDo) (int, error)
	FindByID(int) (model.ToDo, error)
	FindAll() ([]model.ToDo, error)
	DeleteById(int) error
	UpdateById(int, model.ToDo) error
}
