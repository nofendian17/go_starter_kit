package middleware

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a sample HTTP request
	reqBody := []byte(`{"key": "value"}`)
	req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a mock handler to pass to the middleware
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(`{"status": "success"}`))
		if err != nil {
			return
		}
	})

	// Wrap the mock handler with the Logging middleware
	handler := Logging(mockHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the logged output
	expectedLog := `{"status": "success"}`
	assert.Contains(t, rr.Body.String(), expectedLog, "Logged output should contain expected log message")

}
