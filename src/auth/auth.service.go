package auth

import (
	. "flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *UserRepository
}

func (authService *AuthService) Register(user User) (User, *Error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, RegistrationFailed()
	}
	user.Password = hashedPassword

	id, err := authService.UserRepository.create(user)
	if err != nil {
		return User{}, RegistrationFailed()
	}

	created, err := authService.UserRepository.findById(id)
	if err != nil {
		return User{}, UserDoesntExist()
	}

	return created, nil
}

func (authService *AuthService) SignIn(email string, password string) (string, *Error) {
	user, err := authService.UserRepository.findByEmail(email)
	if err != nil {
		return "", InvalidCredentials()
	}
	isPasswordValid := CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		return "", InvalidCredentials()
	}
	return generateToken(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
