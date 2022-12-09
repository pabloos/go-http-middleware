package http_middlware

import (
	"net/http"
)

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
