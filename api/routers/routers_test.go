package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for Routers
func TestUseMw(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	handler := UseMw(mockHandler)
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Hello, World!", resp.Body.String())
}
