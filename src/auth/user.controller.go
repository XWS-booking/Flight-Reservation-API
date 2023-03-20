package auth

import (
	"flight_reservation_api/src/auth/middlewares"
	"flight_reservation_api/src/shared"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type AuthController struct {
	UserService *UserService
}

func (authController *AuthController) constructor(router *mux.Router) {
	adminRouter := router.PathPrefix("").Subrouter()
	adminRouter.Use(middlewares.ExampleMiddleware)
	router.HandleFunc("/register", authController.create).Methods("POST")
	adminRouter.HandleFunc("/user/{id}", authController.findById).Methods("GET")

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
