package database

import (
	"context"
	"log"
	"sync"
	"time"

	"guardian-grid-api/internal/platform/security"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

func InitMongoDB() *mongo.Client {
	mongoOnce.Do(func() {
		uri := security.GetMongoURI()
		
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("failed to connect to MongoDB: %v", err)
		}

		// Ping the primary
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("failed to ping MongoDB: %v", err)
		}

		mongoClient = client
		log.Println("🍃 Connected to MongoDB for telemetry data")
	})

	return mongoClient
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		return InitMongoDB()
	}
	return mongoClient
}

func GetMongoDatabase() *mongo.Database {
	client := GetMongoClient()
	dbName := security.GetMongoDBName()
	return client.Database(dbName)
}

// Deprecated: use GetMongoDatabase().Collection(name) instead
func GetTelemetryCollection() *mongo.Collection {
	return GetMongoDatabase().Collection("telemetry")
}
