package middlewares

import (
	"flight_reservation_api/src/auth/model"
	. "flight_reservation_api/src/shared"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func RolesMiddleware(roles []model.UserRole) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var token *jwt.Token = context.Get(r, "Token").(*jwt.Token)
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				Unauthorized(w)
				return
			}
			role := claims["role"].(float64)
			if !containsRole(roles, int(role)) {
				Forbidden(w)
				return
			}
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func containsRole(roles []model.UserRole, role int) bool {
	for _, r := range roles {
		if int(r) == role {
			return true
		}
	}
	return false
}
