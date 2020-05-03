package entities

import "time"

type UserOtp struct {
	UserId    string    `bson:"userId" json:"userId"`
	UserPhone int       `bson:"userPhone" json:"userPhone"`
	Otp       int       `bson:"otp" json:"otp"`
	IsValid   bool      `bson:"isValid" json:"isValid"`
	Expiry    time.Time `bson:"expiry" json:"expiry"`
	AddedOn   time.Time `bson:"addedOn" json:"addedOn"`
	UpdatedOn time.Time `bson:"updatedOn" json:"updatedOn"`
}
