package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type StartConfig struct {
	Database string
	Host     string
	Password string
	Port     string
	Username string
}

var todoCollection *mongo.Collection

func generateMongoUri(host string, port string) string {
	return fmt.Sprintf("mongodb://%s:%s/", host, port)
}

func StartDatabase(config *StartConfig) {
	clientUri := generateMongoUri(config.Host, config.Port)
	clientCredentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}
	clientOptions := options.Client().ApplyURI(clientUri).SetAuth(clientCredentials)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	todoCollection = client.Database(config.Database).Collection("todos")
}
