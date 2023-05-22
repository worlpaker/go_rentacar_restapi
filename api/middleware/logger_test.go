package middleware

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerMiddleware(t *testing.T) {
	test_data := []struct {
		method       string
		path         string
		statusCode   int
		expected_log string
	}{
		{
			method:       "GET",
			path:         "/",
			statusCode:   http.StatusOK,
			expected_log: "GET / 200",
		},
		{
			method:       "POST",
			path:         "/",
			statusCode:   http.StatusCreated,
			expected_log: "POST / 201",
		},
		{
			method:       "POST",
			path:         "/",
			statusCode:   http.StatusBadRequest,
			expected_log: "POST / 400",
		},
	}

	// Create a recorder to capture the logs
	logOutput := bytes.NewBufferString("")
	log.SetOutput(logOutput)
	for _, k := range test_data {
		// Create a test handler that writes a response with the specified status code
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(k.statusCode)
		})
		// Create a new request
		req, err := http.NewRequest(k.method, k.path, nil)
		assert.Nil(t, err)
		// Create an instance of the middleware and wrap the test handler
		_ = executeRequest(req, Logger(handler))
		// Check if the log output contains the expected log message
		actual_log := logOutput.String()
		if !containsLogMessage(actual_log, k.expected_log) {
			t.Errorf("Expected log message '%s' not found in log output: %s", k.expected_log, actual_log)
		}
	}
}

// containsLogMessage is helper function to check if a log message exists in the log output
func containsLogMessage(logOutput string, logMessage string) bool {
	return bytes.Contains([]byte(logOutput), []byte(logMessage))
}
