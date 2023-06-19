package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRole int64

const (
	ADMINISTRATOR     UserRole = 0
	REGULAR           UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)

type User struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name             string             `bson:"name" json:"name"`
	Surname          string             `bson:"surname" json:"surname"`
	Email            string             `bson:"email" json:"email"`
	Password         string             `bson:"password" json:"password"`
	Role             UserRole           `bson:"role" json:"role"`
	ApiKey           string             `bson:"api_key" json:"api_key"`
	ApiKeyExpiration time.Time          `bson:"api_key_expiration" json:"api_key_expiration"`
	ApiKeyPermanent  bool               `bson:"api_key_permanent" json:"api_key_permanent"`
}

func (user *User) AddApiKey(apiKey string, permanent bool) {
	user.ApiKey = apiKey
	user.ApiKeyPermanent = permanent
	if !permanent {
		currentTime := time.Now()
		user.ApiKeyExpiration = currentTime.AddDate(0, 0, 1)
	}
}
