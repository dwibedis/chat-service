package repository

import (
	"context"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const CollectionUserOtp = "user_otp"
type UserOtp struct {
	collection *mongo.Collection
}

func NewUserOtp(client *mongodb.DBClient) *UserOtp {
	return &UserOtp{collection:client.Collection(CollectionUserOtp)}
}

func (u *UserOtp) SaveUserPhoneAndOtp(ctx context.Context, user *entities.User, otp int) error {
	var userOtp entities.UserOtp
	userOtp.Otp = otp
	userOtp.UserId = user.Id
	userOtp.AddedOn = time.Now()
	userOtp.Expiry = time.Now().Add(time.Minute * 5)
	_, err := u.collection.InsertOne(ctx, &userOtp)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserOtp) GetOtpByPhone(ctx context.Context, phone int) int {
	var userOtp *entities.UserOtp
	filter := bson.M{"userPhone" : phone, "isValid" : true}
	err := u.collection.FindOne(ctx, filter).Decode(&userOtp)
	if err!= nil {
		return 0
	}
	return userOtp.Otp
}

func (u *UserOtp) MarkPhoneOtpInvalid(ctx context.Context, phone int) {
	filter := bson.M{"userPhone" : phone, "isValid" : true}
	update := bson.D{
		{"$set", bson.D{
			{"isValid", false},
			{"updatedOn", time.Now()},
		}},
	}
	_, err := u.collection.UpdateMany(ctx,filter, update)
	if err!= nil {
		//TODO :- set in redis or cache or queue and retry
		log.Println("Error while marking invalid ", err)
		return
	}
	return
}

