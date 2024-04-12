package rest

import (
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/handler"
	mockHealthChecHandler "github.com/nofendian17/gostarterkit/internal/mocks/delivery/rest/handler/healthcheck"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_routes(t *testing.T) {
	mockHandler := &mockHealthChecHandler.Handler{}
	mockHandler.On("Health", mock.Anything).Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	}).Once()
	mockHandler.On("Readiness", mock.Anything).Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	mockHandler.On("Ping", mock.Anything).Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	// Create a new instance of handler.Handler
	h := &handler.Handler{
		Health: mockHandler,
	}

	// You may need to replace this with a mock or real implementation
	// Create a new instance of Server with the mock handler
	s := &server{
		router:  http.NewServeMux(),
		handler: h,
	}

	// Call the routes method to register routes
	s.routes()

	// Create a test server
	testServer := httptest.NewServer(s.router)
	defer testServer.Close()

	// Make GET requests to the registered routes and check if the status codes are OK
	healthResp, err := http.Get(testServer.URL + "/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, healthResp.StatusCode)

	readinessResp, err := http.Get(testServer.URL + "/readiness")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, readinessResp.StatusCode)

	pingResp, err := http.Get(testServer.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, pingResp.StatusCode)
}
