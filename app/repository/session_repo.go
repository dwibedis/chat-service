package repository

import "go.mongodb.org/mongo-driver/mongo"

type Session struct {
	collection *mongo.Collection
}
