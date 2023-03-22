package flights

import (
	"context"
	. "flight_reservation_api/src/flights/model"
	"log"
	"os"
	"time"

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

func (userRepository *FlightRepository) GetAll(date time.Time, startLocation string, endLocation string, seats int) ([]Flight, error) {
	collection := userRepository.getCollection("flights")
	var flights []Flight
	filter := bson.M{"date": primitive.NewDateTimeFromTime(date), "start_location": startLocation, "end_location": endLocation, "seats": seats}
	cur, err := collection.Find(context.TODO(), filter)
	for cur.Next(context.TODO()) {
		var elem Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		flights = append(flights, elem)
	}
	if err != nil {
		return flights, err
	}
	return flights, nil
}

func (flightRepository *FlightRepository) getCollection(key string) *mongo.Collection {
	return flightRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
