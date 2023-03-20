package database

import (
	"context"
	. "fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UniqueField struct {
	Collection string
	Fields     []string
}

func InitDB() (*mongo.Client, error) {
	connection_string := os.Getenv("DATABASE_CONNECTION_STRING")
	Println(connection_string)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(connection_string).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, err
}

func DeclareUnique(client *mongo.Client, fields []UniqueField) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	db := getDatabase(client)
	for _, field := range fields {
		collection := db.Collection(field.Collection)
		for _, fieldname := range field.Fields {
			mod := mongo.IndexModel{
				Keys: bson.M{
					fieldname: 1, // index in descending order
				},
				Options: options.Index().SetUnique(true),
			}
			collection.Indexes().CreateOne(ctx, mod)
		}

	}
}

func getDatabase(db *mongo.Client) *mongo.Database {
	return db.Database(os.Getenv("DATABASE_NAME"))
}
