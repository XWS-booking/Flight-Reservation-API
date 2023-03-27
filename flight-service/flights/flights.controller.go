package flights

import (
	"flight_reservation_api/flights/dtos"
	. "flight_reservation_api/flights/model"
	. "flight_reservation_api/shared"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateFlightController(router *mux.Router, flightService *FlightService) *FlightController {
	controller := &FlightController{FlightService: flightService}
	controller.constructor(router)
	return controller
}

type FlightController struct {
	FlightService *FlightService
}

func (flightController *FlightController) constructor(router *mux.Router) {
	router.HandleFunc("/add", flightController.Create).Methods("POST")
	router.HandleFunc("/getAll/{pageNumber}/{pageSize}", flightController.GetAll).Methods("POST")
	router.HandleFunc("/{id}", flightController.FindById).Methods("GET")
	router.HandleFunc("/{id}", flightController.Delete).Methods("DELETE")
	router.HandleFunc("/{id}/buy-tickets/{quantity}", flightController.BuyTickets).Methods("POST")
	router.HandleFunc("/tickets/listing", flightController.ListTickets).Methods("GET")
}

func (flightController *FlightController) Create(resp http.ResponseWriter, req *http.Request) {
	var flight Flight
	err := DecodeBody(req, &flight)
	fmt.Println(err)
	if err != nil {
		BadRequest(resp, "Something wrong with the data")
		return
	}

	id, e := flightController.FlightService.Create(flight)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, id)
}

func (flightController *FlightController) GetAll(resp http.ResponseWriter, req *http.Request) {
	pageNumber, _ := strconv.Atoi(GetPathParam(req, "pageNumber"))
	pageSize, _ := strconv.Atoi(GetPathParam(req, "pageSize"))
	var flight Flight
	err := DecodeBody(req, &flight)
	if err != nil {
		BadRequest(resp, "Something wrong with the data")
		return
	}

	flights, totalCount, e := flightController.FlightService.GetAll(dtos.NewPageDto(pageNumber, pageSize), flight)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(&resp, dtos.NewFlightPageDto(flights, totalCount))
}

func (flightController *FlightController) FindById(resp http.ResponseWriter, req *http.Request) {
	id := GetPathParam(req, "id")
	flight, e := flightController.FlightService.FindById(StringToObjectId(id))
	if e != nil {
		BadRequest(resp, e.Message)
	}
	Ok(&resp, dtos.NewFlightDto(flight))
}

func (flightController *FlightController) Delete(resp http.ResponseWriter, req *http.Request) {
	id := GetPathParam(req, "id")
	e := flightController.FlightService.Delete(StringToObjectId(id))
	if e != nil {
		BadRequest(resp, e.Message)
	}
	Ok(&resp, e)
}

func (flightController *FlightController) ListTickets(resp http.ResponseWriter, req *http.Request) {
	buyerId := StringToObjectId("6418a6c8e509fcd8c71a4f79")

	tickets, err := flightController.FlightService.FindTicketsByBuyer(buyerId)
	if err != nil {
		BadRequest(resp, "Ticket service unavailable!")
		return
	}

	Ok(&resp, dtos.NewFlightTicketDto(tickets))
}

func (flightController *FlightController) BuyTickets(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("hit")
	flightId := GetPathParam(req, "id")
	quantity := GetPathParam(req, "quantity")
	buyerId := StringToObjectId("6418a6c8e509fcd8c71a4f79")

	quantityNum, err := strconv.Atoi(quantity)
	if err != nil {
		BadRequest(resp, "Quantity should be a number")
		return
	}

	ticketDto := dtos.NewBuyTicketDto(StringToObjectId(flightId), buyerId, quantityNum)
	if err != nil {
		BadRequest(resp, "You request contains wrong data!")
		return
	}
	fmt.Println("conversion done")

	ticketIds, error := flightController.FlightService.BuyTickets(*ticketDto)
	if error != nil {
		BadRequest(resp, error.Message)
		return
	}
	fmt.Println("Service done")
	dto := dtos.NewTicketIdsDto(ticketIds)
	fmt.Println(dto)
	Ok(&resp, dto)
}
