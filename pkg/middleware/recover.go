package middleware

import (
	"fmt"
	"net/http"
)

// recoverErrorHandlerFunc is recover error handle func
type recoverErrorHandlerFunc func(err error, w http.ResponseWriter)

// RecoverMiddleware is middleware config struct
type RecoverMiddleware struct {
	errorHandler recoverErrorHandlerFunc
}

// NewRecoverMiddleware returns new RecoverMiddleware
func NewRecoverMiddleware(e recoverErrorHandlerFunc) *RecoverMiddleware {
	return &RecoverMiddleware{
		errorHandler: e,
	}
}

// Middleware returns panic handle middleware
func (re *RecoverMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(r *RecoverMiddleware) {
			if e := recover(); e != nil {
				err, ok := e.(error)
				if !ok {
					err = fmt.Errorf("%v", e)
				}
				re.errorHandler(err, w)
			}
		}(re)
		next.ServeHTTP(w, r)
	})
}
