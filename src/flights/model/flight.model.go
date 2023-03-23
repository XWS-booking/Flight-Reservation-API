package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
<<<<<<< HEAD
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Seats       int                `bson:"seats" json:"seats"`
	Date        time.Time          `bson:"date" json:"date"`
	Departure   string             `bson:"departure" json:"departure"`
	Destination string             `bson:"destination" json:"destination"`
	Price       float64            `bson:"price" json:"price"`
=======
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Seats         int                `bson:"seats" json:"seats"`
	Date          time.Time          `bson:"date" json:"date"`
	StartLocation string             `bson:"start_location" json:"startLocation"`
	EndLocation   string             `bson:"end_location" json:"endLocation"`
	Price         float64            `bson:"price" json:"price"`
	FreeSeats     int                `bson:"freeSeats" json:"freeSeats"`
}

func (flight *Flight) HasEnoughSeats(seats int) bool {
	return seats >= 1 && seats <= flight.FreeSeats
}

func (flight *Flight) TakeSeats(seats int) {
	flight.FreeSeats = flight.FreeSeats - seats
>>>>>>> Added ticker repository and method for creating tickets
}
