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
	authRouter := router.PathPrefix("").Subrouter()
	authRouter.Use(middlewares.ExampleMiddleware)
	authRouter.Use(middlewares.TokenValidationMiddleware)
	router.HandleFunc("/signup", authController.Signin).Methods("POST")
	router.HandleFunc("/register", authController.create).Methods("POST")
	authRouter.HandleFunc("/user/{id}", authController.findById).Methods("GET")

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

func (authController *AuthController) Signin(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing"))
		return
	}
	if _, ok := r.Header["Password"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Password Missing"))
		return
	}
	valid := authController.UserService.validateUser(r.Header["Email"][0], r.Header["Password"][0])
	if !valid {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Password"))
		return
	}
	tokenString, err := authController.getSignedToken()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}
