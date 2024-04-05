package healthcheck

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	mockUsecase "github.com/nofendian17/gostarterkit/internal/mocks/usecase/healthcheck"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler_Ping(t *testing.T) {

	tests := []struct {
		name string
		want int
	}{
		{
			name: "ping",
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := &mockUsecase.UseCase{}

			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			req.Header.Set("Content-Type", "Application/json")
			rec := httptest.NewRecorder()
			defer rec.Flush()

			h := &handler{
				config:  config.New(),
				useCase: mockUseCase,
			}

			h.Ping().ServeHTTP(rec, req)
			assert.Equal(t, tt.want, rec.Code)
		})
	}
}
