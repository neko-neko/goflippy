package goflippy

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/neko-neko/goflippy/log"
)

var (
	// Version of goflippy
	Version string

	// Hash of goflippy commit hash
	Hash string
)

// HTTP Headers
const (
	HTTPHeaderOrigin = "Origin"

	// Custom headers
	HTTPHeaderXApiKey = "X-API-Key"
)

// HTTPHandlerFunc is custom http handler func
type HTTPHandlerFunc func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)

// HTTPErrorHandlerFunc is http error handler func
type HTTPErrorHandlerFunc func(code int, err error, w http.ResponseWriter)

// handler is base handler
type handler struct {
	// handler is request handler func
	handler HTTPHandlerFunc

	// errorHandler is server error handler func
	errorHandler HTTPErrorHandlerFunc
}

// Handler returns handle func wrapper
func Handler(h HTTPHandlerFunc, eh HTTPErrorHandlerFunc) http.Handler {
	return handler{
		handler:      h,
		errorHandler: eh,
	}
}

// ServeHTTP is wrap handle func executer
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	code, res, err := h.handler(w, r)
	end := time.Now()

	level.Info(log.Logger).Log("message", fmt.Sprintf(
		"[%s] %s %s %s %-7s %s %d %d %s %s %s %s",
		end.Format(time.RFC3339),
		r.RemoteAddr,
		r.Header.Get("X-Forwarded-For"),
		r.Host,
		r.Method,
		r.RequestURI,
		r.ContentLength,
		code,
		strconv.FormatInt(r.ContentLength, 10),
		r.Referer(),
		r.UserAgent(),
		end.Sub(start).String(),
	))
	if err != nil {
		h.errorHandler(code, err, w)
		return
	}

	h.writeJSON(code, res, w)
}

// writeJSON response of JSON format
func (h handler) writeJSON(code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(data)
}
