package main

import (
	. "flight_reservation_api/database"
	. "flight_reservation_api/flights"
	. "flight_reservation_api/flights/repositories/flight"
	. "flight_reservation_api/flights/repositories/tickets"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
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

	flightRepository := &FlightRepository{DB: db, Logger: logger}
	ticketRepository := &TicketRepository{DB: db}
	flightService := &FlightService{FlightRepository: flightRepository, TicketRepository: ticketRepository}
	CreateFlightController(router, flightService)

	startServer(router)
}

func startServer(router *mux.Router) {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	server := http.Server{
		Addr:         ":8080",
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
