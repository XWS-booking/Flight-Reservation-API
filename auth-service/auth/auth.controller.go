package auth

import (
	"flight_reservation_api/auth/dtos"
	"flight_reservation_api/auth/middlewares"
	. "flight_reservation_api/auth/model"
	. "flight_reservation_api/shared"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
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
	router.HandleFunc("/signin", authController.Signin).Methods("POST")
	getUserRouter := router.PathPrefix("").Subrouter()
	getUserRouter.Use(middlewares.TokenValidationMiddleware)
	getUserRouter.Use(middlewares.UserMiddleware)
	getUserRouter.HandleFunc("/user", authController.GetCurrentUser).Methods("GET")
	router.HandleFunc("/register", authController.Register).Methods("POST")
}

func (authController *AuthController) Signin(resp http.ResponseWriter, req *http.Request) {
	var user User
	err := DecodeBody(req, &user)
	if err != nil {
		BadRequest(resp, "Something wrong with the data")
		return
	}

	token, e := authController.AuthService.SignIn(user.Email, user.Password)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, dtos.NewJwtDto(token))
}

func (authController *AuthController) GetCurrentUser(resp http.ResponseWriter, req *http.Request) {
	userid := context.Get(req, "id").(string)
	fmt.Println(userid)

	user, e := authController.AuthService.GetCurrentUser(userid)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, dtos.NewUserDto(user))
}

func (authController *AuthController) Register(resp http.ResponseWriter, req *http.Request) {
	var user User
	err := DecodeBody(req, &user)
	if err != nil {
		BadRequest(resp, "Something wrong with provided data!")
		return
	}
	user.Role = UserRole(REGULAR)
	user, e := authController.AuthService.Register(user)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, dtos.NewUserDto(user))
}
