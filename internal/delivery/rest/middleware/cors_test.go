package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCors(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	corsHandler := Cors(handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://example.com")

	rec := httptest.NewRecorder()

	// Send the request through the CORS middleware
	corsHandler.ServeHTTP(rec, req)

	// Assert that the response code is OK
	assert.Equal(t, http.StatusOK, rec.Code)
}
