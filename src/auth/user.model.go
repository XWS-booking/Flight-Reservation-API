package auth

import (
	. "flight_reservation_api/src/shared"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole int64

const (
	ADMINISTRATOR     UserRole = 0
	REGULAR           UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Surname  string             `bson:"surname" json:"surname"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Role     UserRole           `bson:"role" json:"role"`
}

func (user *User) ValidatePassword(password string) *Error {
	if user.Password != password {
		return InvalidCredentials()
	}
	return nil
}
