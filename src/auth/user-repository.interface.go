package auth

import (
	. "flight_reservation_api/src/auth/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	Create(user User) (primitive.ObjectID, error)
	FindById(id primitive.ObjectID) (User, error)
	FindByApiKey(apiKey string) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) (User, error)
}
