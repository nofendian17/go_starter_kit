package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Stack generates a middleware composed of the provided Middleware functions.
func Stack(mws ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			h = mws[i](h)
		}
		return h
	}
}
