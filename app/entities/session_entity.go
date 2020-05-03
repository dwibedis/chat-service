package entities

type Session struct {
	Id string `bson:"_id" json:"_id"`
	UserId string `bson:"userId" json:"userId"`
	LoggedInAt string `bson:"userId" json:"userId"`

}