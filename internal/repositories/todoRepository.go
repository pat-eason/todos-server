package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/pateason/todo-server/internal/database"
	"github.com/pateason/todo-server/internal/database/entities"
	"github.com/pateason/todo-server/internal/database/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoRecord struct {
	Title     string     `bson:"title"`
	IsActive  bool       `bson:"isActive"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
}

type TodoEntityUpdateModel struct {
	Title string `bson:"title"`
}

var ctx = context.Background()

func RetrieveTodoEntitiesByDate(date time.Time) ([]*entities.TodoEntity, error) {
	output := make([]*entities.TodoEntity, 0)

	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)

	fmt.Println(startDate)
	fmt.Println(endDate)

	createdAtFilter := bson.M{"$gt": startDate, "$lt": endDate}
	filters := bson.M{"createdAt": createdAtFilter}
	cursor, err := database.TodoCollection.Find(ctx, filters)
	if err != nil {
		return output, err
	}

	for cursor.Next(ctx) {
		var todoRecord entities.TodoEntity
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

func CreateTodoEntity(title string) (*entities.TodoEntity, error) {
	insertRecord := TodoRecord{
		Title:     title,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	result, err := database.TodoCollection.InsertOne(ctx, insertRecord)
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

func RetrieveTodoEntity(id string) (*entities.TodoEntity, error) {
	objectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	var result entities.TodoEntity
	filter := bson.M{"_id": objectId}
	err = database.TodoCollection.FindOne(ctx, filter).Decode(&result)
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
	_, err = database.TodoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodoEntityStatus(id string, isActive bool) (*entities.TodoEntity, error) {
	objectId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"isActive": isActive}
	database.TodoCollection.FindOneAndUpdate(ctx, filter, update)

	entity, err := RetrieveTodoEntity(id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
