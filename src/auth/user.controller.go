package auth

import (
	. "flight_reservation_api/src/auth/middlewares"
	"flight_reservation_api/src/shared"
	. "flight_reservation_api/src/shared"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthController struct {
	UserService *UserService
}

func (authController *AuthController) constructor(router *mux.Router) {
	router.HandleFunc("/register", ExampleMiddleware(authController.create)).Methods("POST")
	router.HandleFunc("/user/{id}", authController.findById).Methods("GET")
}

func (authController *AuthController) create(resp http.ResponseWriter, req *http.Request) {
	var user User
	err := DecodeBody(req, &user)
	if err != nil {
		BadRequest(resp, nil)
		return
	}
	id := authController.UserService.create(user)
	Ok(resp, id)
}

func (authController *AuthController) findById(resp http.ResponseWriter, req *http.Request) {
	id := GetPathParam(req, "id")
	objectId := shared.StringToObjectId(id)
	user := authController.UserService.findById(objectId)
	Ok(resp, user)
}
