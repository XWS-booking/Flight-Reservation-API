package dtos

import (
	"flight_reservation_api/src/flights/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightDto struct {
	Id          primitive.ObjectID `json:"id"`
	Seats       int                `json:"seats"`
	Date        time.Time          `json:"date"`
	Departure   string             `json:"departure"`
	Destination string             `json:"destination"`
	Price       float64            `json:"price"`
	FreeSeats   int                `json:"freeSeats"`
}

func NewFlightDto(flight model.Flight) *FlightDto {
	return &FlightDto{
		Id:          flight.Id,
		Seats:       flight.Seats,
		Date:        flight.Date,
		Departure:   flight.Departure,
		Destination: flight.Destination,
		Price:       flight.Price,
		FreeSeats:   flight.FreeSeats,
	}
}
