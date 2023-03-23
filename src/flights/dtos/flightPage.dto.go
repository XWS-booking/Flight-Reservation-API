package dtos

import (
	"flight_reservation_api/src/flights/model"
)

type FlightPageDto struct {
	Data       []model.Flight `json:"data"`
	TotalCount int            `json:"totalCount"`
}

func NewFlightPageDto(flights []model.Flight, totalCount int) *FlightPageDto {
	return &FlightPageDto{
		Data:       flights,
		TotalCount: totalCount,
	}
}
