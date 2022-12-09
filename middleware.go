package http_middlware

import "net/http"

type Middleware func(http.Handler) http.Handler
