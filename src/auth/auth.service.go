package auth

import (
	. "flight_reservation_api/src/shared"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepository *UserRepository
}

func (authService *AuthService) SignIn(email string, password string) (string, *Error) {
	user, err := authService.UserRepository.findByEmail(email)
	if err != nil {
		return "", InvalidCredentials()
	}
	hashedPassword := hashPassword(password)

	e := user.ValidatePassword(hashedPassword)
	if e != nil {
		return "", e
	}

	return generateToken(user)
}

func hashPassword(password string) string {
	return password
}

func generateToken(user User) (string, *Error) {
	var secretKey = []byte(os.Getenv("secret"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["role"] = user.Role
	claims["id"] = user.Id
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", TokenGenerationFailed()
	}

	return tokenString, nil
}
