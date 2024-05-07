package main

import (
	"fmt"
	"github.com/pateason/todo-server/internal/constants"
	"github.com/pateason/todo-server/internal/database"
	"github.com/pateason/todo-server/internal/router"
)

func main() {
	fmt.Println("Connecting to database")
	database.StartDatabase(&database.StartConfig{
		Database: constants.MongoDatabase,
		Host:     constants.MongoHost,
		Password: constants.MongoPass,
		Port:     constants.MongoPort,
		Username: constants.MongoUser,
	})

	fmt.Println("Starting router")
	router.StartRouter()
}
