package auth

import (
	"flight_reservation_api/src/auth/dtos"
	"flight_reservation_api/src/auth/middlewares"
	. "flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"strconv"

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
	loggedUserRouter := authRouter.PathPrefix("").Subrouter()
	loggedUserRouter.Use(middlewares.TokenValidationMiddleware)
	loggedUserRouter.Use(middlewares.UserMiddleware)
	loggedUserRouter.HandleFunc("/user", authController.GetCurrentUser).Methods("GET")
	loggedUserRouter.HandleFunc("/user/api-key/{permanent}", authController.GenerateApiKey).Methods("POST")
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

func (authController *AuthController) GenerateApiKey(resp http.ResponseWriter, req *http.Request) {
	userid := context.Get(req, "id").(string)
	permanent, err := strconv.ParseBool(GetPathParam(req, "permanent"))
	if err != nil {
		BadRequest(resp, "Problem with data!")
		return
	}
	fmt.Println(permanent)
	apiKey, e := authController.AuthService.GenerateApiKey(userid, permanent)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, apiKey)
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
