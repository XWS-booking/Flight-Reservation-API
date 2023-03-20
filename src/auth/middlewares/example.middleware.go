package middlewares

import (
	"log"
	"net/http"
)

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println(h.URL.Path)
		next.ServeHTTP(rw, h)
	})
}
