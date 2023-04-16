package middlewares

import (
	"net/http"

	. "flight_reservation_api/src/shared"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var token *jwt.Token = context.Get(r, "Token").(*jwt.Token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			Unauthorized(rw)
			return
		}
		id := claims["id"].(string)
		context.Set(r, "id", id)
		next.ServeHTTP(rw, r)
	})
}
