package middleware

import (
	"net/http"
	"strings"
)

// MethodRestricted is a middleware that restricts the allowed HTTP methods for a given handler.
func MethodRestricted(next http.HandlerFunc, allowedMethods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !contains(allowedMethods, r.Method) {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// contains checks if a string value exists in a slice of strings, ignoring case sensitivity.
func contains(allowedMethods []string, reqMethod string) (ok bool) {
	for _, item := range allowedMethods {
		if strings.ToUpper(item) == reqMethod {
			ok = true
			return
		}
	}
	return
}
