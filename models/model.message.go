package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message is the model for a message.
type Message struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Message   string             `json:"message,omitempty" bson:"message,omitempty"`
	Timestamp int64              `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Expires   int64              `json:"expires,omitempty" bson:"expires,omitempty"`
	Views     int                `json:"views,omitempty" bson:"views,omitempty"`
	UUID      string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
}
