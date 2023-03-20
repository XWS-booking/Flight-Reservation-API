package auth

import (
	. "flight_reservation_api/src/auth/dtos"
	"flight_reservation_api/src/auth/middlewares"
	"flight_reservation_api/src/shared"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateAuthController(router *mux.Router, userService *UserService, authService *AuthService) *AuthController {
	controller := &AuthController{UserService: userService, AuthService: authService}
	controller.constructor(router)
	return controller
}

type AuthController struct {
	UserService *UserService
	AuthService *AuthService
}

func (authController *AuthController) constructor(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(middlewares.ExampleMiddleware)
	// authRouter.Use(middlewares.TokenValidationMiddleware)
	authRouter.HandleFunc("/signin", authController.Signin).Methods("POST")
	authRouter.HandleFunc("/register", authController.create).Methods("POST")
	authRouter.HandleFunc("/user/{id}", authController.findById).Methods("GET")
	authRouter.Use(middlewares.ErrorHandlerMiddleware)
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

func (authController *AuthController) getSignedToken() (string, error) {
	claimsMap := map[string]string{
		"aud": "frontend.knowsearch.ml",
		"iss": "knowsearch.ml",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := "Secure_Random_String"
	header := "HS256"
	tokenString, err := shared.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
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
	return

}
