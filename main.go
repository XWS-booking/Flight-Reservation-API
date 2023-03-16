package main

import (
	. "flight_reservation_api/src/auth"
	"flight_reservation_api/src/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadEnvs()
	database, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter().StrictSlash(true)

	userControllerFactory := UserControllerFactory{}
	userRepository := &UserRepository{DB: database}
	userService := &UserService{UserRepository: userRepository}
	userControllerFactory.Create(router, userService)

	startServer(router)
}

func startServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(":8000", router))
}
