package flights

import (
	. "flight_reservation_api/src/flights/model"
	"flight_reservation_api/src/shared"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepository *FlightRepository
}

func (flightService *FlightService) Create(flight Flight) (primitive.ObjectID, *shared.Error) {
	created, error := flightService.FlightRepository.Create(flight)
	if error != nil {
		return created, shared.FlightNotCreated()
	}
	return created, nil
}
