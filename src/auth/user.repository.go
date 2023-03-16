package auth

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Client
}

func (userRepository *UserRepository) create(user User) primitive.ObjectID {
	collection := userRepository.getCollection("users")
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID.(primitive.ObjectID)
}

func (userRepository *UserRepository) findById(id primitive.ObjectID) User {
	collection := userRepository.getCollection("users")
	var user User
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (userRepository *UserRepository) getCollection(key string) *mongo.Collection {
	return userRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
