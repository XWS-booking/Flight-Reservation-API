package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Seats         int                `bson:"seats" json:"seats"`
	Date          primitive.DateTime `bson:"date" json:"date"`
	StartLocation string             `bson:"start_location" json:"startLocation"`
	EndLocation   string             `bson:"end_location" json:"endLocation"`
	Price         float64            `bson:"price" json:"price"`
}
