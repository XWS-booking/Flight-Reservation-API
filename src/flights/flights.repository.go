package flights

import (
	"context"
	. "flight_reservation_api/src/flights/model"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (flightRepository *FlightRepository) Create(flight Flight) (primitive.ObjectID, error) {
	collection := flightRepository.getCollection("flights")
	res, err := collection.InsertOne(context.TODO(), flight)
	if err != nil {
		flightRepository.Logger.Println(err)
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (flightRepository *FlightRepository) findById(id primitive.ObjectID) (Flight, error) {
	collection := flightRepository.getCollection("flights")
	var flight Flight
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&flight)
	if err != nil {
		flightRepository.Logger.Println(err)
		return Flight{}, err
	}
	return flight, nil
}

func (flightRepository *FlightRepository) delete(id primitive.ObjectID) error {
	collection := flightRepository.getCollection("flights")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		flightRepository.Logger.Println(err)
		return err
	}
	return nil
}

func (flightRepository *FlightRepository) getCollection(key string) *mongo.Collection {
	return flightRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
