package http_middlware

import "net/http"

// Meet checks preconditions, if not it returns 412
var Meet = BuildValidation(http.StatusPreconditionFailed)

// Limit checks rate limit conditions, if not it returns 429
var Limit = BuildValidation(http.StatusTooManyRequests)

// Validate checks input request (body, headers...), if not it returns 404
var Validate = BuildValidation(http.StatusBadRequest)

// Authenticate checks auth access, if not it returns 401
var Authenticate = BuildValidation(http.StatusUnauthorized)

// Authorization checks role access, if not it returns 403
var Authorization = BuildValidation(http.StatusForbidden)
