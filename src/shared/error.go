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

func UserDoesntExist() *Error {
	return &Error{Message: "User doesn't exist!"}
}
