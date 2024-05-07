package constants

import "os"

var (
	MongoDatabase = os.Getenv("MONGO_DATABASE")
	MongoHost     = os.Getenv("MONGO_HOST")
	MongoPass     = os.Getenv("MONGO_PASS")
	MongoPort     = os.Getenv("MONGO_PORT")
	MongoUser     = os.Getenv("MONGO_USER")
)
