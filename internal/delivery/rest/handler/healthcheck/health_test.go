package healthcheck

import (
	"errors"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	mockUsecase "github.com/nofendian17/gostarterkit/internal/mocks/usecase/healthcheck"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Health(t *testing.T) {
	tests := []struct {
		name     string
		response []interface{}
		want     int
	}{
		{
			name: "should return 200",
			response: []interface{}{
				&response.HealthResponse{
					Version: "1.0",
					Uptime:  "xxx",
					CPU:     "xxx",
					Memory:  "xxx",
				}, nil,
			},
			want: http.StatusOK,
		},
		{
			name: "should return 500",
			response: []interface{}{
				nil, errors.New("error"),
			},
			want: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mockUsecase.UseCase{}
			mockUseCase.On("Health", mock.Anything).Return(tt.response...)

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			req.Header.Set("Content-Type", "Application/json")
			rec := httptest.NewRecorder()
			defer rec.Flush()

			h := &handler{
				config:  config.New(),
				useCase: mockUseCase,
			}

			h.Health().ServeHTTP(rec, req)
			assert.Equal(t, tt.want, rec.Code)
		})
	}
}
