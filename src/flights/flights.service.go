package flights

import (
	"flight_reservation_api/src/flights/dtos"
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

func (flightService *FlightService) GetAll(page dtos.PageDto, flight Flight) ([]Flight, int, *shared.Error) {
	flights, totalCount, err := flightService.FlightRepository.GetAll(page, flight)
	if err != nil {
		return flights, totalCount, shared.FlightsReadFailed()
	}
	return flights, totalCount, nil
}

func (flightService *FlightService) FindById(id primitive.ObjectID) (Flight, *shared.Error) {
	flight, error := flightService.FlightRepository.findById(id)
	if error != nil {
		return flight, shared.FlightNotFound()
	}
	return flight, nil
}

func (FlightService *FlightService) Delete(id primitive.ObjectID) *shared.Error {
	_, e := FlightService.FindById(id)
	if e != nil {
		return shared.FlightNotFound()
	}
	error := FlightService.FlightRepository.delete(id)
	if error != nil {
		return shared.FlightNotDeleted()
	}
	return nil
}
