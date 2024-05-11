package services

import (
	"time"

	"github.com/pateason/todo-server/internal/database/entities"
	"github.com/pateason/todo-server/internal/repositories"
)

type RetrieveTodosModel struct {
	Date *time.Time
}

type CreateTodoModel struct {
	Title string
}

func RetrieveTodos(model RetrieveTodosModel) ([]*entities.TodoEntity, error) {
	var searchDate time.Time
	if model.Date == nil {
		searchDate = time.Now()
	} else {
		searchDate = *model.Date
	}

	entities, err := repositories.RetrieveTodoEntitiesByDate(searchDate)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func RetrieveTodo(id string) (*entities.TodoEntity, error) {
	entity, err := repositories.RetrieveTodoEntity(id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func CreateTodo(model CreateTodoModel) (*entities.TodoEntity, error) {
	entity, err := repositories.CreateTodoEntity(model.Title)
	if err != nil {
		return nil, err
	}
	return entity, err
}

func DeleteTodo(id string) error {
	err := repositories.DeleteTodoEntity(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodoStatus(id string, isActive bool) (*entities.TodoEntity, error) {
	record, err := repositories.UpdateTodoEntityStatus(id, isActive)
	if err != nil {
		return nil, err
	}
	return record, nil
}
