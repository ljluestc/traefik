package traefik_errors

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestTraefikErrorsMiddleware(t *testing.T) {
	testCases := []struct {
		desc            string
		statusCode      int
		backendAttempted bool
		expectedStatus  int
		expectedBody    string
	}{
		{
			desc:            "local 404 error handled",
			statusCode:      http.StatusNotFound,
			backendAttempted: false,
			expectedStatus:  http.StatusNotFound,
			expectedBody:    "Custom error page for 404",
		},
		{
			desc:            "backend 404 error not handled",
			statusCode:      http.StatusNotFound,
			backendAttempted: true,
			expectedStatus:  http.StatusNotFound,
			expectedBody:    "Not Found",
		},
		{
			desc:            "local 502 error handled",
			statusCode:      http.StatusBadGateway,
			backendAttempted: false,
			expectedStatus:  http.StatusBadGateway,
			expectedBody:    "Custom error page for 502",
		},
		{
			desc:            "local 503 error handled",
			statusCode:      http.StatusServiceUnavailable,
			backendAttempted: false,
			expectedStatus:  http.StatusServiceUnavailable,
			expectedBody:    "Custom error page for 503",
		},
		{
			desc:            "non-error response passed through",
			statusCode:      http.StatusOK,
			backendAttempted: true,
			expectedStatus:  http.StatusOK,
			expectedBody:    "OK",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			// Create a custom error handler that returns specific messages for each status code
			errorHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				statusStr := req.Header.Get("X-Traefik-Error-Status")
				if statusStr == "" {
					statusStr = "unknown"
				}
				rw.WriteHeader(test.statusCode)
				rw.Write([]byte("Custom error page for " + statusStr))
			})

			// Create configuration
			config := dynamic.TraefikErrors{
				Status:  []string{"404", "500-599"},
				Service: "error-service",
				Query:   "/error/{status}.html",
			}

			// Create test handler that returns the specified status code
			nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				recorder := rw.(*responseRecorder)
				recorder.SetBackendAttempted(test.backendAttempted)
				rw.WriteHeader(test.statusCode)
				rw.Write([]byte(http.StatusText(test.statusCode)))
			})

			// Create the middleware
			handler, err := New(context.Background(), nextHandler, config, errorHandler, "test")
			require.NoError(t, err)

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
			
			// Create recorder to capture the response
			recorder := httptest.NewRecorder()

			// Call the middleware
			handler.ServeHTTP(recorder, req)

			// Check response
			assert.Equal(t, test.expectedStatus, recorder.Code)
			assert.Contains(t, recorder.Body.String(), test.expectedBody)
		})
	}
}

func TestStatusMatches(t *testing.T) {
	testCases := []struct {
		desc        string
		statusCode  int
		statusList  []string
		shouldMatch bool
	}{
		{
			desc:        "exact match",
			statusCode:  404,
			statusList:  []string{"404"},
			shouldMatch: true,
		},
		{
			desc:        "range match - low end",
			statusCode:  500,
			statusList:  []string{"500-599"},
			shouldMatch: true,
		},
		{
			desc:        "range match - high end",
			statusCode:  599,
			statusList:  []string{"500-599"},
			shouldMatch: true,
		},
		{
			desc:        "range match - middle",
			statusCode:  502,
			statusList:  []string{"500-599"},
			shouldMatch: true,
		},
		{
			desc:        "no match",
			statusCode:  200,
			statusList:  []string{"400-499", "500-599"},
			shouldMatch: false,
		},
		{
			desc:        "multi-range match",
			statusCode:  403,
			statusList:  []string{"400-499", "500-599"},
			shouldMatch: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result := StatusMatches(test.statusCode, test.statusList)
			assert.Equal(t, test.shouldMatch, result)
		})
	}
}

func TestIsTraefikGeneratedError(t *testing.T) {
	testCases := []struct {
		desc            string
		statusCode      int
		backendAttempted bool
		expectedResult  bool
	}{
		{
			desc:            "404 without backend - is Traefik error",
			statusCode:      http.StatusNotFound,
			backendAttempted: false,
			expectedResult:  true,
		},
		{
			desc:            "404 with backend - not a Traefik error",
			statusCode:      http.StatusNotFound,
			backendAttempted: true,
			expectedResult:  false,
		},
		{
			desc:            "502 - always a Traefik error",
			statusCode:      http.StatusBadGateway,
			backendAttempted: true, // doesn't matter
			expectedResult:  true,
		},
		{
			desc:            "503 - always a Traefik error",
			statusCode:      http.StatusServiceUnavailable,
			backendAttempted: true, // doesn't matter
			expectedResult:  true,
		},
		{
			desc:            "504 - always a Traefik error",
			statusCode:      http.StatusGatewayTimeout,
			backendAttempted: true, // doesn't matter
			expectedResult:  true,
		},
		{
			desc:            "500 - not a Traefik error",
			statusCode:      http.StatusInternalServerError,
			backendAttempted: true,
			expectedResult:  false,
		},
		{
			desc:            "200 - not an error",
			statusCode:      http.StatusOK,
			backendAttempted: true,
			expectedResult:  false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result := IsTraefikGeneratedError(test.statusCode, test.backendAttempted)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
