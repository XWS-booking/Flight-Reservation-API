package auth

import (
	"crypto/sha512"
	"encoding/hex"
	. "flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xellDart/uuidapikey"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	UserRepository IUserRepository
}

func (authService *AuthService) Register(user User) (User, *Error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, RegistrationFailed()
	}
	user.Password = hashedPassword

	id, err := authService.UserRepository.Create(user)
	if err != nil {
		return User{}, RegistrationFailed()
	}

	created, err := authService.UserRepository.FindById(id)
	if err != nil {
		return User{}, UserDoesntExist()
	}
	return created, nil
}

func (authService *AuthService) SignIn(email string, password string) (string, *Error) {
	user, err := authService.UserRepository.FindByEmail(email)
	if err != nil {
		return "", InvalidCredentials()
	}
	isPasswordValid := CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		return "", InvalidCredentials()
	}
	return generateToken(user)
}

func (authService *AuthService) GetCurrentUser(userId string) (User, *Error) {
	user, err := authService.UserRepository.FindById(StringToObjectId(userId))
	if err != nil {
		return user, InvalidCredentials()
	}
	return user, nil
}

func (authService *AuthService) GenerateApiKey(userId string, permanent bool) (string, *Error) {
	user, err := authService.UserRepository.FindById(StringToObjectId(userId))
	if err != nil {
		return "", InvalidCredentials()
	}
	apiKey, _ := uuidapikey.Create()
	hashedApiKey := HashData(apiKey)
	user.AddApiKey(hashedApiKey, permanent)
	_, err = authService.UserRepository.Update(user)
	if err != nil {
		return "", ApiKeyGenerationFailed()
	}
	return apiKey, nil
}

func (authService *AuthService) GetUserByApiKey(apiKey string) (User, *Error) {
	var user User
	hashedApiKey := HashData(apiKey)
	user, err := authService.UserRepository.FindByApiKey(hashedApiKey)
	if err != nil {
		return user, InvalidApiKey()
	}
	currentTime := time.Now()
	if user.ApiKeyExpiration.Before(currentTime) && !user.ApiKeyPermanent {
		return user, ApiKeyExpired()
	}
	return user, nil
}

func HashData(data string) string {
	hasher := sha512.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
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
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()
	claims["role"] = user.Role
	claims["id"] = user.Id
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", TokenGenerationFailed()
	}

	return tokenString, nil
}
