package tickets

import (
	"flight_reservation_api/src/flights/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITicketRepository interface {
	CreateMany(tickets []model.Ticket) ([]primitive.ObjectID, error)
	FindAllByBuyer(buyerId primitive.ObjectID) ([]model.Ticket, error)
}
