package tickets

import (
	"context"
	"flight_reservation_api/src/flights/model"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TicketRepository struct {
	DB *mongo.Client
}

func (ticketRepository *TicketRepository) CreateMany(tickets []model.Ticket) ([]primitive.ObjectID, error) {
	collection := ticketRepository.getCollection("tickets")
	res, err := collection.InsertMany(context.TODO(), processTickets(tickets))

	if err != nil {
		return []primitive.ObjectID{}, err
	}

	ids := make([]primitive.ObjectID, len(res.InsertedIDs))
	for i, id := range res.InsertedIDs {
		ids[i] = id.(primitive.ObjectID)
	}

	return ids, nil

}

func (ticketRepository *TicketRepository) DeleteByFlight(flightId primitive.ObjectID) error {
	collection := ticketRepository.getCollection("tickets")
	filter := bson.M{"flightId": flightId}
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}

func (ticketRepository *TicketRepository) FindAllByBuyer(buyerId primitive.ObjectID) ([]model.FlightTicket, error) {
	collection := ticketRepository.getCollection("tickets")
	matchStage := bson.D{{"$match", bson.D{{"buyer", buyerId}}}}
	lookupStage := bson.D{
		{"$lookup", bson.D{
			{"from", "flights"},
			{"localField", "flightId"},
			{"foreignField", "_id"},
			{"as", "flight"},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"ticket", "$$ROOT"},
			{"flight", bson.D{
				{"$arrayElemAt", []interface{}{"$flight", 0}},
			}},
		}},
	}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, projectStage}, options.Aggregate().SetMaxTime(10*time.Second))
	if err != nil {
		return nil, err
	}

	flightTickets, err := loadFlightTickets(cursor, ctx)
	fmt.Println(flightTickets)
	if err != nil {
		return nil, err
	}

	return flightTickets, nil
}

func (ticketRepository *TicketRepository) getCollection(key string) *mongo.Collection {
	return ticketRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}

func processTickets(tickets []model.Ticket) []interface{} {
	var interfaces []interface{} = make([]interface{}, len(tickets))
	for i, v := range tickets {
		interfaces[i] = v
	}
	return interfaces
}

func loadFlightTickets(cursor *mongo.Cursor, ctx context.Context) ([]model.FlightTicket, error) {
	var ticketFlights []model.FlightTicket
	for cursor.Next(ctx) {
		var result model.FlightTicket
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		ticketFlights = append(ticketFlights, result)
	}
	return ticketFlights, nil
}
