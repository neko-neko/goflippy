package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CORSErrorHandlerFunc is error handler func
type CORSErrorHandlerFunc func(code int, err error, w http.ResponseWriter)

// CORSMiddleware is middleware to implement CORS preflight
type CORSMiddleware struct {
	allowAccessOrigins []string
	errorHandler       CORSErrorHandlerFunc
}

// NewCORSMiddleware returns new CorsMiddleware
func NewCORSMiddleware(allowOrigins []string, errorHandler CORSErrorHandlerFunc) *CORSMiddleware {
	return &CORSMiddleware{
		allowAccessOrigins: allowOrigins,
		errorHandler:       errorHandler,
	}
}

// Middleware returns panic handle middleware
func (c *CORSMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		if !c.validateOrigin(origin) {
			c.errorHandler(http.StatusForbidden, fmt.Errorf("%s is invalid origin", origin), w)
			return
		}
		// Do not check request allow methods and headers
		// Because goflippy always allows all methods and headers

		w.Header().Set("Access-Control-Allow-Origin", origin)
		if c.isPreflightRequest(r) {
			c.preflightHandler(w, strings.Trim(r.Header.Get("Access-Control-Request-Headers"), " "))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// validateOrigin is verify valid origin
func (c *CORSMiddleware) validateOrigin(requestOrigin string) bool {
	for _, allowOrigin := range c.allowAccessOrigins {
		if allowOrigin == "*" {
			return true
		}
		if allowOrigin == requestOrigin {
			return true
		}
	}

	return false
}

// isPreflightRequest is check a request is preflight
func (c *CORSMiddleware) isPreflightRequest(r *http.Request) bool {
	return r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != ""
}

func (c *CORSMiddleware) preflightHandler(w http.ResponseWriter, requestHeaders string) {
	w.Header().Set("Access-Control-Max-Age", strconv.Itoa(86400))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "HEAD,GET,POST,PUT,PATCH,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", requestHeaders)

	w.WriteHeader(http.StatusOK)
}
