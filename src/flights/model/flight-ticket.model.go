package model

type FlightTicket struct {
	Ticket Ticket `bson:"ticket,omitempty" json:"ticket"`
	Flight Flight `bson:"flight,omitempty" json:"flight"`
}
