package auth

import (
	. "flight_reservation_api/src/shared"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepository *UserRepository
}

func (userService *UserService) create(user User) primitive.ObjectID {
	return userService.UserRepository.create(user)
}

func (userService *UserService) findById(id primitive.ObjectID) User {
	return userService.UserRepository.findById(id)
}

func (userService *UserService) getOneByEmail(email string) (User, *Error) {
	user, err := userService.UserRepository.findByEmail(email)
	if err != nil {
		return user, UserDoesntExist()
	}
	return user, nil
}
