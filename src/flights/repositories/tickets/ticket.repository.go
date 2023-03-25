package tickets

import (
	"context"
	"flight_reservation_api/src/flights/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
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

func (ticketRepository *TicketRepository) FindAllByBuyer(buyerId primitive.ObjectID) ([]model.Ticket, error) {
	collection := ticketRepository.getCollection("tickets")
	filter := bson.M{"buyerId": buyerId}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return []model.Ticket{}, err
	}

	var result []model.Ticket
	for cur.Next(context.TODO()) {
		var ticket model.Ticket
		err := cur.Decode(&ticket)
		if err != nil {
			return []model.Ticket{}, err
		}
		result = append(result, ticket)
	}
	return result, nil
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
