package auth

import (
	. "flight_reservation_api/src/auth/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	create(user User) primitive.ObjectID
	findById(id primitive.ObjectID) User
	getOneByEmail(email string) User
}
