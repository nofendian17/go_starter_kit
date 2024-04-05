package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func Cors(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		MaxAge:           60,   // 1 minute
		AllowCredentials: true, // Enable CORS with credentials
		Debug:            false,
	}).Handler(h)
}
