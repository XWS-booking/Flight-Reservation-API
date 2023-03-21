package flights

import (
	"context"
	. "flight_reservation_api/src/flights/model"
	"log"
	"os"

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

func (flightRepository *FlightRepository) getCollection(key string) *mongo.Collection {
	return flightRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
