package flights

import (
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	"flight_reservation_api/src/flights/repositories/flight"
	"flight_reservation_api/src/flights/repositories/tickets"
	"flight_reservation_api/src/shared"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepository flight.IFlightRepository
	TicketRepository tickets.ITicketRepository
}

func (flightService *FlightService) Create(flight Flight) (primitive.ObjectID, *shared.Error) {
	created, error := flightService.FlightRepository.Create(flight)
	if error != nil {
		return created, shared.FlightNotCreated()
	}
	return created, nil
}

func (flightService *FlightService) GetAll(page dtos.PageDto, flight Flight) ([]Flight, int, *shared.Error) {
	flights, totalCount, err := flightService.FlightRepository.FindAll(page, flight)
	if err != nil {
		return flights, totalCount, shared.FlightsReadFailed()
	}
	return flights, totalCount, nil
}

func (flightService *FlightService) FindById(id primitive.ObjectID) (Flight, *shared.Error) {
	flight, error := flightService.FlightRepository.FindById(id)
	if error != nil {
		return flight, shared.FlightNotFound()
	}
	return flight, nil
}

func (flightService *FlightService) Delete(id primitive.ObjectID) *shared.Error {
	_, e := flightService.FindById(id)
	if e != nil {
		return shared.FlightNotFound()
	}
	error := flightService.FlightRepository.Delete(id)
	if error != nil {
		return shared.FlightNotDeleted()
	}
	return nil
}

func (flightService *FlightService) BuyTickets(dto dtos.BuyTicketDto) ([]primitive.ObjectID, *shared.Error) {
	flight, err := flightService.FlightRepository.FindById(dto.FlightId)
	if err != nil {
		return []primitive.ObjectID{}, shared.FlightNotFound()
	}

	hasSeats := flight.HasEnoughSeats(dto.Quantity)
	if !hasSeats {
		return []primitive.ObjectID{}, shared.NotEnoughSeats()
	}
	tickets := createTickets(dto)
	ticketIds, err := flightService.TicketRepository.CreateMany(tickets)

	flight.TakeSeats(dto.Quantity)
	flightService.FlightRepository.Update(&flight)

	if err != nil {
		return []primitive.ObjectID{}, shared.TicketServiceUnavailable()
	}
	return ticketIds, nil
}

func (flightService *FlightService) FindTicketsByBuyer(userId primitive.ObjectID) ([]FlightTicket, *shared.Error) {
	tickets, err := flightService.TicketRepository.FindAllByBuyer(userId)
	if err != nil {
		return []FlightTicket{}, shared.TicketServiceUnavailable()
	}
	return tickets, nil
}

func createTickets(dto dtos.BuyTicketDto) []Ticket {
	tickets := make([]Ticket, 0)

	for _ = range make([]struct{}, dto.Quantity) {
		tickets = append(tickets, Ticket{BuyerId: dto.BuyerId, FlightId: dto.FlightId})
	}
	return tickets
}
