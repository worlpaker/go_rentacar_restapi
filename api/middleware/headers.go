package middleware

import "net/http"

// SetHeaders is a middleware that sets custom headers to the response.
func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add your headers here
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
