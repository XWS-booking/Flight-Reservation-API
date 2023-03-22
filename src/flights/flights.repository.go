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

func (flightRepository *FlightRepository) GetAll(pageNumber int, pageSize int, date primitive.DateTime, startLocation string, endLocation string, seats int) ([]Flight, int, error) {
	collection := flightRepository.getCollection("flights")
	var flights []Flight
	filter := bson.D{}
	toDate := date.Time().AddDate(0, 0, 1)
	if date.Time().IsZero() {
		toDate = time.Now().AddDate(100, 0, 0)
	}
	filter = bson.D{{Key: "end_location", Value: bson.D{{Key: "$regex", Value: "(?i).*" + endLocation + ".*"}}},
		{Key: "start_location", Value: bson.D{{Key: "$regex", Value: "(?i).*" + startLocation + ".*"}}},
		{Key: "seats", Value: bson.D{{Key: "$gte", Value: seats}}},
		{Key: "date", Value: bson.D{{Key: "$gte", Value: date}, {Key: "$lt", Value: toDate}}}}
	options := new(options.FindOptions)
	options.SetSkip(int64((pageNumber - 1) * pageSize))
	options.SetLimit(int64(pageSize))

	cur, err := collection.Find(context.TODO(), filter, options)
	totalCount, _ := flightRepository.GetTotalCount(filter)
	if err != nil {
		return flights, totalCount, err
	}

	for cur.Next(context.TODO()) {
		var elem Flight
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
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

func (flightRepository *FlightRepository) getCollection(key string) *mongo.Collection {
	return flightRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}