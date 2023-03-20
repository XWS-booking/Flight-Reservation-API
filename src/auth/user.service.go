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

func (userService *UserService) getOneByEmail(email string) User {
	return userService.UserRepository.getOneByEmail(email)
}

func (userService *UserService) validateUser(email string, password string) bool {
	usr := userService.getOneByEmail(email)
	return usr.ValidatePassword(password)
}
