package shared

type Error struct {
	Message string
}

func InvalidCredentials() *Error {
	return &Error{Message: "Invalid credentials!"}
}

func TokenGenerationFailed() *Error {
	return &Error{Message: "Token generation failed!"}
}

func TokenValidationFailed() *Error {
	return &Error{Message: "Token validation failed!"}
}

func UserDoesntExist() *Error {
	return &Error{Message: "User doesn't exist!"}
}

func RegistrationFailed() *Error {
	return &Error{Message: "Registration data invalid or user with given email already exists!"}
}

func FlightNotCreated() *Error {
	return &Error{Message: "Flight creation went wrong!"}
}

func FlightsReadFailed() *Error {
	return &Error{Message: "Cannot read flights!"}
}

func FlightsCountFailed() *Error {
	return &Error{Message: "Cannot count flights!"}
}
