package middleware

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodRestrictedMiddleware(t *testing.T) {
	type methods []string
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	allowedMethods := methods{"GET", "POST", "patch"}
	allowedReq, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	disallowedReq, err := http.NewRequest("PUT", "/", nil)
	assert.Nil(t, err)
	lowcaseReq, err := http.NewRequest("PATCH", "/", nil)
	assert.Nil(t, err)
	middleware := MethodRestricted(handler, allowedMethods)
	resp_allowed := executeRequest(allowedReq, middleware)
	assert.Equal(t, http.StatusOK, resp_allowed.Code)
	resp_disallowed := executeRequest(disallowedReq, middleware)
	assert.Equal(t, http.StatusMethodNotAllowed, resp_disallowed.Code)
	resp_lowcase := executeRequest(lowcaseReq, middleware)
	assert.Equal(t, http.StatusOK, resp_lowcase.Code)
}

func BenchmarkContains(b *testing.B) {
	arr := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	val := "GET"

	for i := 0; i < b.N; i++ {
		contains(arr, val)
	}
}
