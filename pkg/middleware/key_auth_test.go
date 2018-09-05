package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/neko-neko/goflippy/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func authSuccessHandler(r *http.Request, id string) *http.Request {
	return httptest.NewRequest("GET", "/test", nil)
}

func authErrorHandler(code int, apiKey string, w http.ResponseWriter) {
	w.WriteHeader(code)
}

func TestKeyAuthMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TEST"))
	}).Methods("GET")

	router.Use(NewKeyAuthMiddleware(
		func(projectID string) (string, error) {
			return "", nil
		}, authSuccessHandler, authErrorHandler).Middleware)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, "TEST", rr.Body.String())
}

func TestKeyAuthMiddlewareWhenError(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TEST"))
	}).Methods("GET")

	router.Use(NewKeyAuthMiddleware(
		func(projectID string) (string, error) {
			return "", errors.New("TEST")
		}, authSuccessHandler, authErrorHandler).Middleware)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}
