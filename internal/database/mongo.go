package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myapp/internal/config"
)

// DB represents the database connection.
var DB *mongo.Client

// ConnectDB establishes a connection to MongoDB.
func ConnectDB(cfg *config.Config) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDBURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	DB = client
	log.Println("Connected to MongoDB successfully!")
}

// GetCollection returns a collection from the database.
func GetCollection(cfg *config.Config, collectionName string) *mongo.Collection {
	return DB.Database(cfg.MongoDBName).Collection(collectionName)
}