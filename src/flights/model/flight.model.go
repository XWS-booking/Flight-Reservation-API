package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Seats       int                `bson:"seats" json:"seats"`
	Date        time.Time          `bson:"date" json:"date"`
	Departure   string             `bson:"departure" json:"departure"`
	Destination string             `bson:"destination" json:"destination"`
	Price       float64            `bson:"price" json:"price"`
}
