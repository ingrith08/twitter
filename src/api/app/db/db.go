package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(uri string, dbName string) *MongoRepository {
	db := connect(uri, dbName)
	return &MongoRepository{db}
}

func connect(uri string, dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetMaxPoolSize(10))
	check(err, "Failed to connect to mongo database")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	check(err, "Failed to connect to mongo database")

	err = client.Ping(ctx, readpref.Primary())
	check(err, "Failed to ping the mongo database")

	db := client.Database(dbName)

	log.Println("Connecting to mongodb", fmt.Sprintf("Connected to %s database", dbName), map[string]string{})

	return db
}

func check(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err.Error(), map[string]string{})
		os.Exit(1)
	}
}
