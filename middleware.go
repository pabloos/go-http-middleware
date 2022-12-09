package http_middlware

import "net/http"

// Middleware wraps handlers to do some pre or post request processing
type Middleware func(http.Handler) http.Handler
