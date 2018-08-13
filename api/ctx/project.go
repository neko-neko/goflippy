package ctx

import (
	"context"
	"net/http"
)

type contextProjectID string

const contextProjectIDKey contextProjectID = "ProjectID"

// GetProjectID retruns ProjectID associated APIKey
func GetProjectID(ctx context.Context) string {
	return ctx.Value(contextProjectIDKey).(string)
}

// CreateRequestWithContext create new context in http.Request with ProjectID
func CreateRequestWithContext(r *http.Request, id string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), contextProjectIDKey, id))
}
