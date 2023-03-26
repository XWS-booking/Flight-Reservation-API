package auth

import (
	"flight_reservation_api/src/auth/dtos"
	"flight_reservation_api/src/auth/middlewares"
	. "flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
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
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signin", authController.Signin).Methods("POST")
	authRouter.HandleFunc("/user", authController.GetCurrentUser).Methods("GET")
	http.Handle("/auth/user", middlewares.TokenValidationMiddleware(authRouter))
	http.Handle("/auth/user", middlewares.UserMiddleware(authRouter))
	authRouter.HandleFunc("/register", authController.Register).Methods("POST")
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

	Ok(&resp, user)
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
