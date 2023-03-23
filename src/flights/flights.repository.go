package flights

import (
	"context"
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (flightRepository *FlightRepository) GetAll(page dtos.PageDto, flight Flight) ([]Flight, int, error) {
	collection := flightRepository.getCollection("flights")
	var flights []Flight
	filter := bson.D{}
	toDate := flight.Date.AddDate(0, 0, 1)
	if flight.Date.IsZero() {
		toDate = time.Now().AddDate(100, 0, 0)
	}
	filter = bson.D{{Key: "end_location", Value: bson.D{{Key: "$regex", Value: "(?i).*" + flight.EndLocation + ".*"}}},
		{Key: "start_location", Value: bson.D{{Key: "$regex", Value: "(?i).*" + flight.StartLocation + ".*"}}},
		{Key: "seats", Value: bson.D{{Key: "$gte", Value: flight.Seats}}},
		{Key: "date", Value: bson.D{{Key: "$gte", Value: flight.Date}, {Key: "$lt", Value: toDate}}}}
	options := new(options.FindOptions)
	options.SetSkip(int64((page.PageNumber - 1) * page.PageSize))
	options.SetLimit(int64(page.PageSize))

	cur, err := collection.Find(context.TODO(), filter, options)
	totalCount, _ := flightRepository.GetTotalCount(filter)
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

func (flightRepository *FlightRepository) GetTotalCount(filter bson.D) (int, error) {
	collection := flightRepository.getCollection("flights")
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return int(count), err
	}
	return int(count), nil
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
