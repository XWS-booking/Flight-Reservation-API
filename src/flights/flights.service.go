package flights

import (
	. "flight_reservation_api/src/flights/model"
	"flight_reservation_api/src/shared"
	"time"

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

func (flightService *FlightService) GetAll(date time.Time, startLocation string, endLocation string, seats int) ([]Flight, *shared.Error) {
	flights, error := flightService.FlightRepository.GetAll(date, startLocation, endLocation, seats)
	if error != nil {
		return flights, shared.FlightsReadFailed()
	}
	return flights, nil
}
