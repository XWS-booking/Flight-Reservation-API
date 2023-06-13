package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BuyerId  primitive.ObjectID `bson:"buyer" json:"buyer"`
	FlightId primitive.ObjectID `bson:"flightId" json:"flightId"`
}
