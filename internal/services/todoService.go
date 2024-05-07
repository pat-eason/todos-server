package services

import (
	"github.com/pateason/todo-server/internal/database"
	"time"
)

type RetrieveTodosModel struct {
	Date time.Time
}

type CreateTodoModel struct {
	Title   string
	Content string
}

func RetrieveTodos(model RetrieveTodosModel) ([]*database.TodoEntity, error) {
	entities, err := database.RetrieveTodoEntities()
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func RetrieveTodo(id string) (*database.TodoEntity, error) {
	entity, err := database.RetrieveTodoEntity(id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func CreateTodo(model CreateTodoModel) (*database.TodoEntity, error) {
	entity, err := database.CreateTodoEntity(model.Title, model.Content)
	if err != nil {
		return nil, err
	}
	return entity, err
}

func DeleteTodo(id string) error {
	err := database.DeleteTodoEntity(id)
	if err != nil {
		return err
	}
	return nil
}
