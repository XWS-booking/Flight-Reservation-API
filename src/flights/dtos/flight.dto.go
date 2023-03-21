package dtos

import (
	"flight_reservation_api/src/flights/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightDto struct {
	Id            primitive.ObjectID `json:"id"`
	Seats         int                `json:"seats"`
	Date          primitive.DateTime `json:"date"`
	StartLocation string             `json:"startLocation"`
	EndLocation   string             `json:"endLocation"`
	Price         float64            `json:"price"`
}

func NewFlightDto(flight model.Flight) *FlightDto {
	return &FlightDto{
		Id:            flight.Id,
		Seats:         flight.Seats,
		Date:          flight.Date,
		StartLocation: flight.StartLocation,
		EndLocation:   flight.EndLocation,
		Price:         flight.Price,
	}
}
