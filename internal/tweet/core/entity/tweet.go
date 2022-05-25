package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	Message string `bson:"message" json:"message"`
}

type SaveTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}

type ResponseTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"userid" json:"userid,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
