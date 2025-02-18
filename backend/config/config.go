package config

import (
	"context"
	"log"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(GetEnv("MONGO_DB_URL"))

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	MongoClient = client
	return MongoClient
}

func GetCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}
