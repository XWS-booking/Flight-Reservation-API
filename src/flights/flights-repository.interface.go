package flights

import (
	. "flight_reservation_api/src/flights/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IFlightRepository interface {
	create(flight Flight) primitive.ObjectID
	findById(id primitive.ObjectID) Flight
	findAll() []Flight
	delete(id primitive.ObjectID) error
	getAll(date time.Time, startLocation string, endLocation string, seats int) []Flight
}
