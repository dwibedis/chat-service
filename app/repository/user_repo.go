package repository

import (
	"context"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/app/util"
	"github.com/dwibedis/chat-service/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const CollectionUser = "users"
type User struct {
	collection *mongo.Collection
}

func NewUserRepo(client *mongodb.DBClient) *User {
	return &User{collection:client.Database.Collection(CollectionUser)}
}

func (u *User) AddUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	userId, err := util.GenerateRandomString(32)
	if err != nil {
		return nil, err
	}
	user.Id = userId
	user.AddedOn = time.Now()
	user.IsValidated = false
	result, err := u.collection.InsertOne(ctx, user);
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(string)
	return user, nil
}

func (u *User) GetUserByPhone(ctx context.Context, phone int) *entities.User {
	var user entities.User
	filter := bson.M{"phone" : phone}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return &user
}

func (u *User) GetUserById(ctx context.Context, userId string) *entities.User  {
	var user entities.User
	filter := bson.M{"_id" : userId}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return &user
}