package database

import (
	"context"
	"github.com/pateason/todo-server/internal/database/utils"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type TodoRecord struct {
	Title     string     `bson:"title"`
	Content   string     `bson:"content"`
	IsActive  bool       `bson:"isActive"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt"`
}

type TodoEntity struct {
	Id        *string    `bson:"_id"`
	Title     string     `bson:"title"`
	Content   string     `bson:"content"`
	IsActive  bool       `bson:"isActive"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt"`
}

var ctx = context.Background()

func RetrieveTodoEntities() ([]*TodoEntity, error) {
	output := make([]*TodoEntity, 0)

	cursor, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {
		return output, err
	}

	for cursor.Next(ctx) {
		var todoRecord TodoEntity
		err := cursor.Decode(&todoRecord)
		if err != nil {
			return output, err
		}
		output = append(output, &todoRecord)
	}

	if err := cursor.Err(); err != nil {
		return output, err
	}

	if err := cursor.Close(ctx); err != nil {
		return output, err
	}

	return output, nil
}

func CreateTodoEntity(title string, content string) (*TodoEntity, error) {
	insertRecord := TodoRecord{
		Title:     title,
		Content:   content,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	result, err := todoCollection.InsertOne(ctx, insertRecord)
	if err != nil {
		return nil, err
	}

	extractedId := utils.ExtractId(result)
	createdEntity, err := RetrieveTodoEntity(extractedId)
	if err != nil {
		return nil, err
	}

	return createdEntity, nil
}

func RetrieveTodoEntity(id string) (*TodoEntity, error) {
	objectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	var result TodoEntity
	filter := bson.M{"_id": objectId}
	err = todoCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteTodoEntity(id string) error {
	objectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	_, err = todoCollection.DeleteOne(ctx, filter)
	return nil
}
