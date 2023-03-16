package auth

import "github.com/gorilla/mux"

type IUserControllerFactory interface {
	create(userService *UserService) AuthController
}

type UserControllerFactory struct{}

func (userControllerFactory *UserControllerFactory) Create(router *mux.Router, userService *UserService) *AuthController {
	controller := &AuthController{UserService: userService}
	controller.constructor(router)
	return controller
}
