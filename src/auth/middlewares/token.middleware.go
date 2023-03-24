package middlewares

import (
	"flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			shared.Unauthorized(rw)
			return
		}
		bearer := strings.Split(r.Header["Authorization"][0], " ")
		token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			var secretKey = []byte(os.Getenv("JWT_SECRET"))
			return secretKey, nil
		})
		if err != nil {
			shared.Unauthorized(rw)
			return
		}
		context.Set(r, "Token", token)
		next.ServeHTTP(rw, r)
	})
}
