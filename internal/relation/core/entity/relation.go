package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Relation struct {
	UserID         string `bson:"userid" json:"userid"`
	UserRelationID string `bson:"userrelationid" json:"userrelationid"`
}

type RequestRelation struct {
	UserRelationID string `json:"userRelationId"`
}

type ListTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userrelationid,omitempty"`
	Tweet          struct {
		Message string             `bson:"message" json:"message,omitempty"`
		Date    time.Time          `bson:"date" json:"date,omitempty"`
		ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	}
}
