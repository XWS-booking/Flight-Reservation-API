package middlewares

import (
	"flight_reservation_api/src/shared"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if token is present
		if r.Header["Authorization"] == nil {
			shared.Unauthorized(rw)
			return
		}
		bearer := strings.Split(r.Header["Authorization"][0], " ")
		token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
			nesto, ok := token.Method.(*jwt.SigningMethodHMAC)
			fmt.Println(nesto)
			fmt.Println(ok)
			if !ok {
				return "", fmt.Errorf("ok")
			}
			fmt.Println("bruh")
			fmt.Println(token)
			return token, nil

		})
		fmt.Println(err)
		if token == nil {
			shared.Unauthorized(rw)
			return
		}
		context.Set(r, "Token", token)
		next.ServeHTTP(rw, r)
	})
}
