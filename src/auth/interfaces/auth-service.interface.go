package interfaces

import (
	. "flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
)

type IAuthService interface {
	Register(user User) (User, *Error)
	SignIn(email string, password string) (string, *Error)
	GetCurrentUser(userId string) (User, *Error)
	GenerateApiKey(userId string, permanent bool) (string, *Error)
	GetUserByApiKey(apiKey string) (User, *Error)
}
