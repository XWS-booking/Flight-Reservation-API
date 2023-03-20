package auth

import (
	. "flight_reservation_api/src/auth/dtos"
	"flight_reservation_api/src/auth/middlewares"
	. "flight_reservation_api/src/shared"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAuthController(router *mux.Router, authService *AuthService) *AuthController {
	controller := &AuthController{AuthService: authService}
	controller.constructor(router)
	return controller
}

type AuthController struct {
	AuthService *AuthService
}

func (authController *AuthController) constructor(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middlewares.ExampleMiddleware)
	// authRouter.Use(middlewares.TokenValidationMiddleware)
	authRouter.HandleFunc("/signin", authController.Signin).Methods("POST")
	authRouter.Use(middlewares.ErrorHandlerMiddleware)
}

func (authController *AuthController) Signin(resp http.ResponseWriter, req *http.Request) {
	var user User
	err := DecodeBody(req, &user)
	if err != nil {
		BadRequest(resp, "Something wrong with the data")
	}

	token, e := authController.AuthService.SignIn(user.Email, user.Password)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	tokenDto := JwtDto{Token: token}
	Ok(resp, tokenDto)
}