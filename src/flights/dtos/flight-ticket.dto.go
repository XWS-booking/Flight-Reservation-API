package dtos

import "flight_reservation_api/src/flights/model"

type FlightTicketDto struct {
	Tickets []model.FlightTicket `json:"tickets"`
}

func NewFlightTicketDto(tickets []model.FlightTicket) *FlightTicketDto {
	return &FlightTicketDto{
		Tickets: tickets,
	}
}
