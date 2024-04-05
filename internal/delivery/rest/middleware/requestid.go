package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"net/http"
)

func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = uuid.New().String()
			r.Header.Set("X-Request-Id", requestID)
		}

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		logger := slog.WithContext(ctx)
		logger.Infof("request started with request_id: %s", requestID)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
