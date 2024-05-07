package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExtractId(id *mongo.InsertOneResult) string {
	return id.InsertedID.(primitive.ObjectID).Hex()
}

func ConvertToObjectId(hexId string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}
