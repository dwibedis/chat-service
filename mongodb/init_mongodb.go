package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoUrl = "mongodb://localhost:27017/merchant_relay_service"
const DbName = "chat-service"

type DBClient struct {
	*mongo.Database
}

func InitMongoDb() (*DBClient, error) {
	ctx := context.Background()
	// Set client options
	clientOptions := options.Client().ApplyURI(MongoUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database(DbName)
	mongoClient := &DBClient{Database: db}
	return mongoClient, nil
}