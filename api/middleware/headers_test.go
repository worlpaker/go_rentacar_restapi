package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func executeRequest(r *http.Request, mw http.Handler) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	mw.ServeHTTP(w, r)
	return w
}

func TestSetHeadersMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", contentType)
	})
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp := executeRequest(req, SetHeaders(handler))
	expected := "application/json"
	actual := resp.Header().Get("Content-Type")
	assert.Equal(t, expected, actual)
}
