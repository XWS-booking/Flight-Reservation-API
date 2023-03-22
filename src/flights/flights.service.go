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

func (flightService *FlightService) GetAll(pageNumber int, pageSize int, date primitive.DateTime, startLocation string, endLocation string, seats int) ([]Flight, int, *shared.Error) {
	flights, totalCount, err := flightService.FlightRepository.GetAll(pageNumber, pageSize, date, startLocation, endLocation, seats)
	if err != nil {
		return flights, totalCount, shared.FlightsReadFailed()
	}
	return flights, totalCount, nil
}
