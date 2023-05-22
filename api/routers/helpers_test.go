package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func executeReqMid(r *http.Request, mw http.HandlerFunc) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	mw.ServeHTTP(w, r)
	return w
}

func TestHandle(t *testing.T) {
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	handler := Handle(mockHandler).Methods("GET")
	resp := executeReqMid(req, handler)
	assert.Equal(t, http.StatusOK, resp.Code)
	req, err = http.NewRequest("POST", "/", nil)
	assert.Nil(t, err)
	resp = executeReqMid(req, handler)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
}

func TestApplyMiddlewares(t *testing.T) {
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	mockMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	handler := ApplyMiddlewares(http.HandlerFunc(mockHandler), mockMiddleware)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
