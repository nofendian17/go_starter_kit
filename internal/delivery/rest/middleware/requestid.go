package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

// RequestIDKey is a custom type for the request ID key in context.
type RequestIDKey string

const (
	// RequestIDHeader is the name of the header used to pass request ID.
	RequestIDHeader = "X-Request-Id"
	// DefaultRequestIDKey is the default value for the request ID key.
	DefaultRequestIDKey RequestIDKey = "requestID"
)

// RequestID middleware generates a unique request ID and adds it to the request context.
func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String()
			r.Header.Set(RequestIDHeader, requestID)
		}

		ctx := context.WithValue(r.Context(), DefaultRequestIDKey, requestID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
