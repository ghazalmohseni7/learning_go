package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db *mongo.Database

func Connect() error {
	err := LoadEnv()
	if err != nil {
		return err
	}
	url := os.Getenv("MONGODB_URL")
	port := os.Getenv("MONGODB_PORT")
	uri := url + ":" + port
	dbName := os.Getenv("DB_NAME")
	fmt.Printf("dburl %s: ", url)
	fmt.Printf("dbport : %s", port)
	fmt.Printf("dbpdbName %s: ", dbName)
	client, errordb := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if errordb != nil {
		return errordb
	}
	db = client.Database(dbName)
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

func CloseConnection() error {
	return db.Client().Disconnect(context.Background())
}
