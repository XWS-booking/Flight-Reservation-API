package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserService struct {
	UserRepository *UserRepository
}

func (userService *UserService) create(user User) primitive.ObjectID {
	return userService.UserRepository.create(user)
}

func (userService *UserService) findById(id primitive.ObjectID) User {
	return userService.UserRepository.findById(id)
}
