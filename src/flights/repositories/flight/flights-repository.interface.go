package flight

import (
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IFlightRepository interface {
	Create(flight Flight) (primitive.ObjectID, error)
	FindById(id primitive.ObjectID) (Flight, error)
	Delete(id primitive.ObjectID) error
	FindAll(page dtos.PageDto, flight Flight) ([]Flight, int, error)
	Update(flight *Flight) error
}