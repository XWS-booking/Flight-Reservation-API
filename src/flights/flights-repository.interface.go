package flights

import (
	. "flight_reservation_api/src/flights/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IFlightRepository interface {
	create(flight Flight) primitive.ObjectID
	findById(id primitive.ObjectID) Flight
	findAll() []Flight
	delete(id primitive.ObjectID) error
	getAll(pageNumber int, pageSize int, date primitive.DateTime, startLocation string, endLocation string, seats int) []Flight
	GetTotalCount(flightRepository *FlightRepository) (int, error)
}
