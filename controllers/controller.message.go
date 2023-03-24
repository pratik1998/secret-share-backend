package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	config "github.com/pratik1998/secret-share-backend/configs"
	models "github.com/pratik1998/secret-share-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Store message in database
func StoreMessage(msg string) string {
	var message models.Message
	message.Message = msg
	message.Timestamp = time.Now().Unix()
	message.Expires = message.Timestamp + 3600
	message.Views = 1
	message.UUID = uuid.New().String()
	collection := config.GetCollection("messages")
	_, err := collection.InsertOne(context.Background(), message)
	if err != nil {
		panic(err)
	}
	return message.UUID
}

// Get message from database
func GetMessage(uuid string) string {
	var message models.Message
	collection := config.GetCollection("messages")
	err := collection.FindOne(context.Background(), models.Message{UUID: uuid}).Decode(&message)
	if err != nil {
		fmt.Println("Unknown UUID:", uuid)
		return "Invalid UUID"
	}

	// Check if message has expired
	if message.Expires < time.Now().Unix() {
		fmt.Println("Message expired:", uuid)
		return "No Longer Available"
	}

	// Update message views
	message.Views++
	filter := models.Message{ID: message.ID}
	// To resolve must use $set error
	update := bson.D{{"$set", bson.D{{"views", message.Views}}}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error updating message views:", uuid)
		return "Error updating message views"
	}

	return message.Message
}
