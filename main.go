package main

import (
	. "flight_reservation_api/src/auth"
	. "flight_reservation_api/src/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	LoadEnvs()
	db, err := InitDB()
	DeclareUnique(db, []UniqueField{
		{Collection: "users", Fields: []string{"email"}},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter().StrictSlash(true)

	userControllerFactory := UserControllerFactory{}
	logger := log.New(os.Stdout, "[Users-api] ", log.LstdFlags)
	userRepository := &UserRepository{DB: db, Logger: logger}
	userService := &UserService{UserRepository: userRepository}
	userControllerFactory.Create(router, userService)

	startServer(router)
}

func startServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(":8000", router))
}
