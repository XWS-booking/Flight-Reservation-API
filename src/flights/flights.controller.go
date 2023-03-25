package flights

import (
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	. "flight_reservation_api/src/shared"
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
	authRouter := router.PathPrefix("/flights").Subrouter()
	authRouter.HandleFunc("/add", flightController.Create).Methods("POST")
	authRouter.HandleFunc("/getAll/{pageNumber}/{pageSize}", flightController.GetAll).Methods("POST")
	authRouter.HandleFunc("/{id}", flightController.FindById).Methods("GET")
	authRouter.HandleFunc("/{id}", flightController.Delete).Methods("DELETE")
	authRouter.HandleFunc("/{id}/buy-tickets/{quantity}", flightController.BuyTickets).Methods("POST")
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

	ticketIds, error := flightController.FlightService.BuyTickets(*ticketDto)
	if error != nil {
		BadRequest(resp, error.Message)
		return
	}
	dto := dtos.NewTicketIdsDto(ticketIds)
	Ok(&resp, dto)
}
