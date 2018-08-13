package middleware

import (
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/neko-neko/goflippy/pkg/handler"
	"github.com/neko-neko/goflippy/pkg/log"
)

// findProjectFunc is find project from datastore func
type findProjectFunc func(projectID string) (string, error)

// authErrorHandlerFunc is key auth error handle func
type authErrorHandlerFunc func(code int, apiKey string, w http.ResponseWriter)

// authSuccessHandlerFunc is keu auth success handle func
type authSuccessHandlerFunc func(r *http.Request, id string) *http.Request

// KeyAuthMiddleware is HTTP header authentication middleware
type KeyAuthMiddleware struct {
	findFunc       findProjectFunc
	successHandler authSuccessHandlerFunc
	errorHandler   authErrorHandlerFunc
}

// NewKeyAuthMiddleware returns new KeyAuthMiddleware
func NewKeyAuthMiddleware(f findProjectFunc, s authSuccessHandlerFunc, e authErrorHandlerFunc) *KeyAuthMiddleware {
	return &KeyAuthMiddleware{
		findFunc:       f,
		successHandler: s,
		errorHandler:   e,
	}
}

// Middleware return key auth middleware
func (k *KeyAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(handler.HTTPHeaderXApiKey)
		level.Debug(log.Logger).Log("message", "request API Key", "key", apiKey)

		projectID, err := k.findFunc(apiKey)
		if err != nil {
			k.errorHandler(http.StatusForbidden, apiKey, w)
			return
		}
		level.Debug(log.Logger).Log("message", "project id", "id", projectID)

		next.ServeHTTP(w, k.successHandler(r, projectID))
	})
}
