package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// Logging logs incoming requests and outgoing responses, and captures request and response bodies
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Capture request body
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Error reading request body:", err)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// Create a buffer to capture response body
		responseBodyBuffer := bytes.NewBuffer(nil)
		captureWriter := &responseCaptureWriter{ResponseWriter: w, buf: responseBodyBuffer}

		// Serve the next handler
		next.ServeHTTP(captureWriter, r)

		// Calculate elapsed time
		elapsed := time.Since(start).Milliseconds()

		slog.With(slog.String("type", "tdr")).
			// With("request_id", r.Context().Value(DefaultRequestIDKey)).
			With(slog.Group("request",
				slog.String("source", r.RemoteAddr),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				// slog.Any("headers", r.Header),
				slog.Any("body", string(requestBody)),
			)).With(slog.Group("response",
			slog.Int("code", captureWriter.Status()),
			slog.String("body", responseBodyBuffer.String()),
			slog.String("elapsed", fmt.Sprintf("%dms", elapsed)),
		)).Info("request complete")
	})
}

// responseCaptureWriter wraps the original http.ResponseWriter to capture response body and status code
type responseCaptureWriter struct {
	http.ResponseWriter
	buf         *bytes.Buffer
	statusCode  int
	wroteHeader bool
}

// WriteHeader captures the status code
func (w *responseCaptureWriter) WriteHeader(code int) {
	if !w.wroteHeader {
		w.statusCode = code
		w.ResponseWriter.WriteHeader(code)
		w.wroteHeader = true
	}
}

// Write captures the response body
func (w *responseCaptureWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

// Status returns the captured status code
func (w *responseCaptureWriter) Status() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}
