package middlewares

import (
	"flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func RolesMiddleware(roles []model.UserRole) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var token *jwt.Token = context.Get(r, "Token").(*jwt.Token)
			claims, ok := token.Claims.(jwt.MapClaims)
			fmt.Println(ok)
			if ok && token.Valid {
				role := claims["role"].(string)
				fmt.Println(role)
			}
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
func DecodeToken(token string) (*model.UserRole, *Error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	// ... error handling
	if err != nil {
		return nil, TokenValidationFailed()
	}
	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	return nil, nil
}
