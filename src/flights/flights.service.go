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

func (flightService *FlightService) GetAll(pageNumber int, pageSize int, date time.Time, startLocation string, endLocation string, seats int) ([]Flight, int, *shared.Error) {
	flights, err := flightService.FlightRepository.GetAll(pageNumber, pageSize, date, startLocation, endLocation, seats)
	totalCount, err2 := flightService.FlightRepository.GetTotalCount()
	if err != nil {
		return flights, totalCount, shared.FlightsReadFailed()
	}
	if err2 != nil {
		return flights, totalCount, shared.FlightsCountFailed()
	}
	return flights, totalCount, nil
}
