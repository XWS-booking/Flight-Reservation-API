package middlewares

import (
	"flight_reservation_api/src/auth/interfaces"
	"github.com/gorilla/context"
	"net/http"
)

func ApiKeyMiddleware(authService interfaces.IAuthService) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, r *http.Request) {
			if r.Header["Api-Key"] == nil {
				context.Set(r, "api-key-authorization", false)
				h.ServeHTTP(rw, r)
				return
			}
			apiKey := r.Header["Api-Key"][0]
			user, err := authService.GetUserByApiKey(apiKey)
			if err != nil {
				context.Set(r, "api-key-authorization", false)
				h.ServeHTTP(rw, r)
				return
			}
			context.Set(r, "id", user.Id.Hex())
			context.Set(r, "role", user.Role)
			context.Set(r, "api-key-authorization", true)
			h.ServeHTTP(rw, r)
		}
		return http.HandlerFunc(fn)
	}
}
