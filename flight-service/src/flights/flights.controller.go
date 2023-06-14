package flights

import (
	"flight_reservation_api/src/auth"
	"flight_reservation_api/src/auth/middlewares"
	"flight_reservation_api/src/auth/model"
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateFlightController(router *mux.Router, flightService *FlightService, authService *auth.AuthService) *FlightController {
	controller := &FlightController{FlightService: flightService, AuthService: authService}
	controller.constructor(router)
	return controller
}

type FlightController struct {
	FlightService *FlightService
	AuthService   *auth.AuthService
}

func (flightController *FlightController) constructor(router *mux.Router) {
	flightRouter := router.PathPrefix("/flights").Subrouter()

	protectedRoute(flightRouter, "/add", "POST", flightController.AuthService, []model.UserRole{model.ADMINISTRATOR}, flightController.Create)
	protectedRoute(flightRouter, "/{id}", "DELETE", flightController.AuthService, []model.UserRole{model.ADMINISTRATOR}, flightController.Delete)
	protectedRoute(flightRouter, "/{id}/buy-tickets/{quantity}", "POST", flightController.AuthService, []model.UserRole{model.REGULAR}, flightController.BuyTickets)
	protectedRoute(flightRouter, "/tickets/listing", "GET", flightController.AuthService, []model.UserRole{model.REGULAR}, flightController.ListTickets)

	flightRouter.HandleFunc("/getAll/{pageNumber}/{pageSize}", flightController.GetAll).Methods("POST")
	flightRouter.HandleFunc("/{id}", flightController.FindById).Methods("GET")
}

func protectedRoute(router *mux.Router, route string, method string, authService *auth.AuthService, roles []model.UserRole, f func(http.ResponseWriter, *http.Request)) {
	newRouter := router.PathPrefix("/").Subrouter()
	newRouter.Use(middlewares.ApiKeyMiddleware(authService))
	newRouter.Use(middlewares.TokenValidationMiddleware)
	newRouter.Use(middlewares.RolesMiddleware(roles))
	newRouter.Use(middlewares.UserMiddleware)
	newRouter.HandleFunc(route, f).Methods(method)
}

func (flightController *FlightController) Create(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("hit")
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
	id := context.Get(req, "id").(string)
	buyerId := StringToObjectId(id)

	tickets, err := flightController.FlightService.FindTicketsByBuyer(buyerId)
	if err != nil {
		BadRequest(resp, "Ticket service unavailable!")
		return
	}

	Ok(&resp, dtos.NewFlightTicketDto(tickets))
}

func (flightController *FlightController) BuyTickets(resp http.ResponseWriter, req *http.Request) {
	flightId := GetPathParam(req, "id")
	quantity := GetPathParam(req, "quantity")
	id := context.Get(req, "id").(string)
	buyerId := StringToObjectId(id)

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
