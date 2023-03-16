package main

import (
	"flight_reservation_api/src/controller"
	"fmt"
)

func main() {

	testVar := controller.TestStruct{
		Field: "Srdjan",
	}

	fmt.Println(testVar)

}
