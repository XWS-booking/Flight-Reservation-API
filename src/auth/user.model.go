package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole int64

const (
	ADMINISTRATOR     UserRole = 0
	REGULAR           UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)

type User struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string
	Surname string
	Role    UserRole
}
