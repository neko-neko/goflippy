package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/neko-neko/goflippy/pkg/log"
)

// HTTP Headers
const (
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

	log.Info(
		"remote-addr", r.RemoteAddr,
		"host", r.Host,
		"method", r.Method,
		"uri", r.RequestURI,
		"content-length", strconv.FormatInt(r.ContentLength, 10),
		"status-code", code,
		"referer", r.Referer(),
		"user-agent", r.UserAgent(),
		"elapsed-time", end.Sub(start).String(),
	)
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
