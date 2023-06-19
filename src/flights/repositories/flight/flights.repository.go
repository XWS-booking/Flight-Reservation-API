package flight

import (
	"context"
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
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

func (flightRepository *FlightRepository) FindAll(page dtos.PageDto, flight Flight) ([]Flight, int, error) {
	collection := flightRepository.getCollection("flights")
	var flights []Flight
	filter := bson.D{}
	toDate := flight.Date.AddDate(0, 0, 1)
	if flight.Date.IsZero() {
		toDate = time.Now().AddDate(100, 0, 0)
	}
	filter = bson.D{{Key: "destination", Value: bson.D{{Key: "$regex", Value: "(?i).*" + flight.Destination + ".*"}}},
		{Key: "departure", Value: bson.D{{Key: "$regex", Value: "(?i).*" + flight.Departure + ".*"}}},
		{Key: "freeSeats", Value: bson.D{{Key: "$gte", Value: flight.FreeSeats}}},
		{Key: "date", Value: bson.D{{Key: "$gte", Value: flight.Date}, {Key: "$lt", Value: toDate}}}}
	options := new(options.FindOptions)
	options.SetSkip(int64((page.PageNumber - 1) * page.PageSize))
	options.SetLimit(int64(page.PageSize))

	cur, err := collection.Find(context.TODO(), filter, options)
	totalCount, _ := flightRepository.getTotalCount(filter)
	if err != nil {
		return flights, totalCount, err
	}

	for cur.Next(context.TODO()) {
		var elem Flight
		err := cur.Decode(&elem)
		if err != nil {
			return flights, totalCount, err
		}
		flights = append(flights, elem)
	}
	return flights, totalCount, nil
}

func (flightRepository *FlightRepository) GetFlightsForReservation(date time.Time, departure string, destination string) ([]Flight, error) {
	collection := flightRepository.getCollection("flights")
	flights := make([]Flight, 0)

	filter := bson.D{
		{Key: "destination", Value: bson.D{{Key: "$regex", Value: "(?i).*" + destination + ".*"}}},
		{Key: "departure", Value: bson.D{{Key: "$regex", Value: "(?i).*" + departure + ".*"}}},
		{Key: "date", Value: bson.D{{Key: "$gte", Value: date}, {Key: "$lt", Value: date.AddDate(0, 0, 1)}}},
		{Key: "freeSeats", Value: bson.D{{Key: "$gt", Value: 0}}},
	}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return flights, err
	}
	for cur.Next(context.TODO()) {
		var elem Flight
		err := cur.Decode(&elem)
		if err != nil {
			return flights, err
		}
		flights = append(flights, elem)
	}
	return flights, nil

}

func (flightRepository *FlightRepository) getTotalCount(filter bson.D) (int, error) {
	collection := flightRepository.getCollection("flights")
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return int(count), err
	}
	return int(count), nil
}

func (flightRepository *FlightRepository) FindById(id primitive.ObjectID) (Flight, error) {
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

func (flightRepository *FlightRepository) Delete(id primitive.ObjectID) error {
	collection := flightRepository.getCollection("flights")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		flightRepository.Logger.Println(err)
		return err
	}
	return nil
}

func (flightRepository *FlightRepository) Update(flight *Flight) error {
	collection := flightRepository.getCollection("flights")
	fmt.Println(flight)
	filter := bson.M{"_id": flight.Id}
	update := bson.D{{"$set", bson.D{{"freeSeats", flight.FreeSeats}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (flightRepository *FlightRepository) getCollection(key string) *mongo.Collection {
	return flightRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
