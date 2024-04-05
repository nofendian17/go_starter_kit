package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// Define some sample middlewares
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do something before calling the next handler
			next.ServeHTTP(w, r)
		})
	}
	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do something before calling the next handler
			next.ServeHTTP(w, r)
		})
	}

	// Create a test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a stack with the sample middlewares
	stack := Stack(middleware1, middleware2)

	// Wrap the handler with the stack
	stackHandler := stack(handler)

	// Create a mock request and recorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Send the request through the stack
	stackHandler.ServeHTTP(rec, req)

	// Assert that the response code is OK
	assert.Equal(t, http.StatusOK, rec.Code)
}
