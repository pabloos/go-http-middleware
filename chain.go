package http_middlware

import (
	"net/http"
)

// Chain links middlewares, returning a new middleware
func Chain(mids ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		if len(mids) < 1 {
			return handler
		}

		for i := len(mids) - 1; i >= 0; i-- {
			handler = mids[i](handler)
		}

		return handler
	}
}
