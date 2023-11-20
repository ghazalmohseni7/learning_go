package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db *mongo.Database

func Connect(state string) error {
	var DBName string
	err := LoadEnv()
	if err != nil {
		return err
	}
	url := os.Getenv("MONGODB_URL")
	port := os.Getenv("MONGODB_PORT")
	uri := url + ":" + port
	if state == "DEV" {
		DBName = os.Getenv("DEV_DB_NAME")
		//fmt.Println("here in the connection db is : ", DBName)
	} else if state == "TEST" {
		DBName = os.Getenv("TEST_DB_NAME")
		//fmt.Println("here in the connection db is : ", DBName)
	}

	//fmt.Printf("dburl %s: ", url)
	//fmt.Printf("dbport : %s", port)
	//fmt.Printf("dbpdbName %s: ", DBName)
	client, errordb := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if errordb != nil {
		return errordb
	}
	db = client.Database(DBName)
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

func CloseConnection() error {
	return db.Client().Disconnect(context.Background())
}

func CreateCollection(collection_name string) bool {
	err := db.CreateCollection(context.TODO(), collection_name)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func EmptyTestDB() {
	//time.Sleep(20 * time.Second)
	//fmt.Println("in the empty db")
	//fmt.Println(db.Name())
	err := db.Drop(context.Background())
	if err != nil {
		fmt.Println("BBBBBBBBBBBBBB")
		fmt.Println(err.Error())
	} else {
		fmt.Println("test db is deleted")
	}

	//collectionNames, errorCollections := db.ListCollectionNames(context.Background(), bson.M{})
	//if errorCollections != nil {
	//	fmt.Println("some error happen \n" + errorCollections.Error())
	//}
	//fmt.Println(collectionNames)
	//for _, coll := range collectionNames {
	//	fmt.Println(coll)
	//	errorDrop := db.Collection(coll).Drop(context.Background())
	//	if errorDrop != nil {
	//		fmt.Println("can not delete the collection\n" + errorDrop.Error())
	//	} else {
	//		fmt.Println("collection " + coll + " is deleted")
	//	}
	//}
}
