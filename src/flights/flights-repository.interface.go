package flights

import (
	. "flight_reservation_api/src/flights/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IFlightRepository interface {
	create(flight Flight) primitive.ObjectID
	findById(id string) Flight
	findAll() []Flight
	delete(id primitive.ObjectID) error
}
