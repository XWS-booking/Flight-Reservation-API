package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type BuyTicketDto struct {
	FlightId primitive.ObjectID `json:"flightId"`
	BuyerId  primitive.ObjectID `json:"buyerId"`
	Quantity int                `json:"quantity"`
}

func NewBuyTicketDto(flightId primitive.ObjectID, buyerId primitive.ObjectID, quantity int) *BuyTicketDto {
	return &BuyTicketDto{
		FlightId: flightId,
		BuyerId:  buyerId,
		Quantity: quantity,
	}
}
