package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoCOnnection() (*MongoDb, error) {
	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	return nil, nil
}

type MongoDb struct {
}
