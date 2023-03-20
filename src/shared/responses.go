package shared

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Ok(resp http.ResponseWriter, payload any) {
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(payload)
}

func BadRequest(resp http.ResponseWriter, message string) {
	resp.WriteHeader(http.StatusBadRequest)
	payload := ErrorResponse{Status: http.StatusBadRequest, Message: message}
	json.NewEncoder(resp).Encode(payload)
}

func NotFound(resp http.ResponseWriter, message string) {
	resp.WriteHeader(http.StatusNotFound)
	payload := ErrorResponse{Status: http.StatusNotFound, Message: message}
	json.NewEncoder(resp).Encode(payload)
}

func Unauthorized(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusUnauthorized)
	payload := ErrorResponse{Status: http.StatusUnauthorized, Message: "Unauthorized"}
	json.NewEncoder(resp).Encode(payload)
}

func Forbidden(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusForbidden)
	payload := ErrorResponse{Status: http.StatusForbidden, Message: "Forbidden"}
	json.NewEncoder(resp).Encode(payload)
}

func JsonResponse(resp http.ResponseWriter, statusCode int, payload interface{}) {
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(payload)
}

func EmptyResponse(resp http.ResponseWriter, statusCode int) {
	resp.WriteHeader(statusCode)
}
