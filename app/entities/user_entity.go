package entities

import "time"

type User struct {
	Id          string    `bson:"_id" json:"id"`
	Phone       int       `bson:"phone" json:"phone"`
	Name        string    `bson:"name" json:"name"`
	IsValidated bool      `bson:"isValid" json:"isValid"`
	AddedOn     time.Time `bson:"addedOn" json:"addedOn"`
	UpdatedOn   time.Time `bson:"updatedOn" json:"updatedOn"`
}
