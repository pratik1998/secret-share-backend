package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(config *Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Database.GetURL()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client Instance
var DB *mongo.Client

func init() {
	config := GetConfig()
	DB = ConnectDB(config)
}

// GetDB returns the database instance
func GetDB() *mongo.Client {
	return DB
}

// GetCollection returns the collection instance
func GetCollection(collection string) *mongo.Collection {
	return DB.Database("secret-share").Collection(collection)
}
