package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestSizeMiddleware(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		w.Write([]byte(string(rune(len(body)))))
	})
	// Test case 1: success
	requestBody := strings.NewReader(strings.Repeat("a", 1024*5)) // 5 KB body
	req := httptest.NewRequest("POST", "/", requestBody)
	handler := RequestSize(1024 * 10)(testHandler)
	resp := executeRequest(req, handler)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test case 2: fails
	// Create a test request with a large body
	requestBody = strings.NewReader(strings.Repeat("a", 1024*20)) // 20 KB body
	req = httptest.NewRequest("POST", "/", requestBody)
	resp = executeRequest(req, handler)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Bad Request\n", resp.Body.String())
}
