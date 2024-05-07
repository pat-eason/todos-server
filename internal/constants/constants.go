package constants

import "os"

var (
	MongoHost = os.Getenv("MONGO_HOST")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPass = os.Getenv("MONGO_PASS")
)
