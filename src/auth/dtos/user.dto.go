package dtos

import (
	"flight_reservation_api/src/auth/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDto struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Surname string             `json:"surname"`
	Email   string             `json:"email"`
	Role    model.UserRole     `json:"role"`
}

func NewUserDto(user model.User) *UserDto {
	return &UserDto{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Role:    user.Role,
	}
}
