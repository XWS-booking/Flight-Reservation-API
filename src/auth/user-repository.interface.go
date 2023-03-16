package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type IUserRepository interface {
	create(user User) primitive.ObjectID
	findById(id primitive.ObjectID) User
}
