package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/neko-neko/goflippy/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func errorHandler(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func TestRecoverMiddleware(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TEST"))
	}).Methods("GET")

	router.Use(NewRecoverMiddleware(errorHandler).Middleware)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, "TEST", rr.Body.String())
}

func TestRecoverMiddlewareWhenError(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		panic("TEST")
	}).Methods("GET")

	router.Use(NewRecoverMiddleware(errorHandler).Middleware)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
