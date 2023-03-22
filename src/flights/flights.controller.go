package flights

import (
	"flight_reservation_api/src/flights/dtos"
	. "flight_reservation_api/src/flights/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	authRouter.HandleFunc("/getAll/{pageNumber}/{pageSize}/{startLocation}/{endLocation}/{seats}/{date}", flightController.GetAll).Methods("GET")
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

	Ok(resp, id)
}

func (flightController *FlightController) GetAll(resp http.ResponseWriter, req *http.Request) {
	pageNumber, _ := strconv.Atoi(GetPathParam(req, "pageNumber"))
	pageSize, _ := strconv.Atoi(GetPathParam(req, "pageSize"))
	startLocation := GetPathParam(req, "startLocation")
	endLocation := GetPathParam(req, "endLocation")
	seats, _ := strconv.Atoi(GetPathParam(req, "seats"))
	dateParam := GetPathParam(req, "date")
	if startLocation == "null" {
		startLocation = ""
	}
	if endLocation == "null" {
		endLocation = ""
	}

	date, err := time.Parse("2006-01-02T15:04:05.000Z", dateParam)
	if err != nil {
		date = time.Time{}
	}

	flights, totalCount, e := flightController.FlightService.GetAll(pageNumber, pageSize, date, startLocation, endLocation, seats)
	if e != nil {
		BadRequest(resp, e.Message)
		return
	}

	Ok(resp, dtos.NewFlightPageDto(flights, totalCount))
}
