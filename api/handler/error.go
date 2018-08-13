package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/neko-neko/goflippy/pkg/log"
)

// authErrorResponse is unauthorize error response structure
type authErrorResponse struct {
	Message string `json:"message"`
}

// AuthErrorHandler is handle server error func
func AuthErrorHandler(code int, apiKey string, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(&authErrorResponse{
		Message: fmt.Sprintf("invalid API Key %s", apiKey),
	})
}

// recoverErrorResponse is internal server error response structure
type recoverErrorResponse struct {
	Message string `json:"message"`
}

// RecoverErrorHandler is handle internal server error func
func RecoverErrorHandler(err error, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(&recoverErrorResponse{
		Message: "internal server error",
	})
	level.Error(log.Logger).Log("message", "server panic", "err", err)
}
