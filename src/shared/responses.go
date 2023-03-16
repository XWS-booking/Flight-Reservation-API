package shared

import (
	"encoding/json"
	"net/http"
)

func Ok(resp http.ResponseWriter, payload any) {
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(payload)
}

func BadRequest(resp http.ResponseWriter, payload any) {
	resp.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(resp).Encode(payload)
}
