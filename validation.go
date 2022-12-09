package http_middlware

import (
	"net/http"
)

// Check es a function
type Check func(req *http.Request) error

// ValidateMiddlware is a middleware built from a lists of checks
type ValidateMiddlware func(checks ...Check) Middleware

// BuildValidation builds a ValidateMiddlware from a http status error
func BuildValidation(statusErr int) ValidateMiddlware {
	return func(checks ...Check) Middleware {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				for _, check := range checks {
					err := check(req)
					if err != nil {
						http.Error(res, err.Error(), statusErr)
						return
					}
				}

				next.ServeHTTP(res, req)
			})
		}
	}
}
