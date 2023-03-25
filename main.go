package main

import (
	. "flight_reservation_api/src/auth"
	. "flight_reservation_api/src/database"
	. "flight_reservation_api/src/flights"
	. "flight_reservation_api/src/flights/repositories/flight"
	"flight_reservation_api/src/flights/repositories/tickets"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
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
	logger := log.New(os.Stdout, "[Users-api] ", log.LstdFlags)
	userRepository := &UserRepository{DB: db, Logger: logger}
	authService := &AuthService{UserRepository: userRepository}
	CreateAuthController(router, authService)

	flightRepository := &FlightRepository{DB: db, Logger: logger}
	ticketRepository := &tickets.TicketRepository{DB: db}
	flightService := &FlightService{FlightRepository: flightRepository, TicketRepository: ticketRepository}
	CreateFlightController(router, flightService)

	startServer(router)
}

func startServer(router *mux.Router) {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	server := http.Server{
		Addr:         ":8000",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
