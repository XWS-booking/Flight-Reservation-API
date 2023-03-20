package middlewares

import (
	"flight_reservation_api/src/shared"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if token is present
		if _, ok := r.Header["Token"]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Missing"))
			return
		}
		token := r.Header["Token"][0]
		check, err := shared.ValidateToken(token, "Secure_Random_String")

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Token Validation Failed"))
			return
		}
		if !check {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Invalid"))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Authorized Token"))

	})
}
