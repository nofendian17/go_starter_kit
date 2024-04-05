package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestID(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(DefaultRequestIDKey)
		if requestID != nil {
			w.Header().Set(RequestIDHeader, requestID.(string))
		}
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Send the request through the middleware
	RequestID(handler).ServeHTTP(rec, req)

	// Assert that the response code is OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that the request ID is set in the response header
	assert.NotEmpty(t, rec.Header().Get(RequestIDHeader))
}
