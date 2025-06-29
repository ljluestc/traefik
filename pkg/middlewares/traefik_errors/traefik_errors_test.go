package traefik_errors

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestTraefikErrors(t *testing.T) {
	// Create temporary error page files
	tempDir := t.TempDir()

	errorPage404 := filepath.Join(tempDir, "404.html")
	errorPage500 := filepath.Join(tempDir, "500.html")

	err := os.WriteFile(errorPage404, []byte("<html><body>Custom 404 error page</body></html>"), 0666)
	require.NoError(t, err)

	err = os.WriteFile(errorPage500, []byte("<html><body>Custom 500 error page</body></html>"), 0666)
	require.NoError(t, err)

	testCases := []struct {
		desc            string
		errorPages      map[string]string
		returnedStatus  int
		isTraefikError  bool
		expectedStatus  int
		expectedContent string
	}{
		{
			desc: "Traefik error with custom page",
			errorPages: map[string]string{
				"404": errorPage404,
			},
			returnedStatus:  http.StatusNotFound,
			isTraefikError:  true,
			expectedStatus:  http.StatusNotFound,
			expectedContent: "<html><body>Custom 404 error page</body></html>",
		},
		{
			desc: "Upstream error without interference",
			errorPages: map[string]string{
				"404": errorPage404,
			},
			returnedStatus:  http.StatusNotFound,
			isTraefikError:  false,
			expectedStatus:  http.StatusNotFound,
			expectedContent: "Upstream 404 content",
		},
		{
			desc: "Traefik error with no matching custom page",
			errorPages: map[string]string{
				"500": errorPage500,
			},
			returnedStatus:  http.StatusNotFound,
			isTraefikError:  true,
			expectedStatus:  http.StatusNotFound,
			expectedContent: "", // Default error handler
		},
		{
			desc:            "No custom error pages configured",
			errorPages:      map[string]string{},
			returnedStatus:  http.StatusNotFound,
			isTraefikError:  true,
			expectedStatus:  http.StatusNotFound,
			expectedContent: "", // Default error handler
		},
		{
			desc: "Success response shouldn't be affected",
			errorPages: map[string]string{
				"404": errorPage404,
			},
			returnedStatus:  http.StatusOK,
			isTraefikError:  true,
			expectedStatus:  http.StatusOK,
			expectedContent: "OK content",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			config := dynamic.TraefikErrors{
				ErrorPages: test.errorPages,
			}

			// Create a test handler
			next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if test.isTraefikError {
					// Simulate a Traefik-generated error
					req = req.WithContext(context.WithValue(req.Context(), "traefikGeneratedError", true))
					req.Header.Set("X-Traefik-Internal-Error", "true")
				}

				if test.returnedStatus == http.StatusOK {
					rw.WriteHeader(http.StatusOK)
					_, _ = rw.Write([]byte("OK content"))
					return
				}

				rw.WriteHeader(test.returnedStatus)
				if !test.isTraefikError {
					_, _ = rw.Write([]byte("Upstream 404 content"))
				}
			})

			handler, err := New(context.Background(), next, config, "test-errors")
			require.NoError(t, err)

			// Create a test request
			req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, req)

			assert.Equal(t, test.expectedStatus, recorder.Code)

			responseBody, err := io.ReadAll(recorder.Body)
			require.NoError(t, err)

			if test.expectedContent != "" {
				assert.Equal(t, test.expectedContent, string(responseBody))
			}
		})
	}
}

func TestTraefikErrorsInvalidConfig(t *testing.T) {
	testCases := []struct {
		desc        string
		errorPages  map[string]string
		expectError bool
	}{
		{
			desc: "Invalid status code",
			errorPages: map[string]string{
				"invalid": "page.html",
			},
			expectError: true,
		},
		{
			desc: "Non-existent error page",
			errorPages: map[string]string{
				"404": "/does/not/exist.html",
			},
			expectError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			config := dynamic.TraefikErrors{
				ErrorPages: test.errorPages,
			}

			_, err := New(context.Background(), http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}), config, "test-errors")
			
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
