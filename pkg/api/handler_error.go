package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/pkg/server/middleware/traefik_errors"
)

// ErrorEntry represents a single error entry in the health API
type ErrorEntry struct {
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Method     string    `json:"method"`
	Host       string    `json:"host"`
	Path       string    `json:"path"`
	IP         string    `json:"ip"`
	Headers    string    `json:"headers"`
	Time       time.Time `json:"time"`
}

// ErrorBuffer stores the most recent errors for the health API
type ErrorBuffer struct {
	sync.Mutex
	errors    []ErrorEntry
	maxErrors int
}

var (
	// Global error buffer that stores the most recent errors
	errorBuffer = &ErrorBuffer{
		errors:    make([]ErrorEntry, 0, 100),
		maxErrors: 100, // Default max errors to store
	}
)

// AddError adds a new error to the buffer
func (eb *ErrorBuffer) AddError(entry ErrorEntry) {
	eb.Lock()
	defer eb.Unlock()

	// Add new error at the beginning
	eb.errors = append([]ErrorEntry{entry}, eb.errors...)

	// Trim if we have too many
	if len(eb.errors) > eb.maxErrors {
		eb.errors = eb.errors[:eb.maxErrors]
	}
}

// GetErrors returns all stored errors
func (eb *ErrorBuffer) GetErrors() []ErrorEntry {
	eb.Lock()
	defer eb.Unlock()

	// Make a copy to avoid race conditions
	result := make([]ErrorEntry, len(eb.errors))
	copy(result, eb.errors)
	return result
}

// RecordError records an error in the global error buffer
func RecordError(details traefik_errors.ErrorDetails) {
	entry := ErrorEntry{
		StatusCode: details.StatusCode,
		Status:     details.Status,
		Method:     details.Method,
		Host:       details.Host,
		Path:       details.Path,
		IP:         details.IP,
		Headers:    details.Headers,
		Time:       details.Time,
	}

	errorBuffer.AddError(entry)
}

// GetHealthErrorsHandler returns a handler that displays recent errors
func GetHealthErrorsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		errors := errorBuffer.GetErrors()
		
		rw.Header().Set("Content-Type", "application/json")
		
		err := json.NewEncoder(rw).Encode(errors)
		if err != nil {
			log.Ctx(req.Context()).Error().Err(err).Msg("Error encoding health errors")
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

// ErrorLoggerMiddleware intercepts errors and records them for the health API
func ErrorLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Create a custom response writer to capture status code
		respRecorder := &statusRecorder{
			ResponseWriter: rw,
			status:         http.StatusOK,
		}

		// Call the next handler
		next.ServeHTTP(respRecorder, req)

		// If the response was an error, log it
		if respRecorder.status >= 400 {
			// Check if we have detailed error info in the context
			if details, ok := req.Context().Value("traefik.error.details").(traefik_errors.ErrorDetails); ok {
				RecordError(details)
			} else {
				// Create basic error entry without detailed info
				RecordError(traefik_errors.ErrorDetails{
					StatusCode: respRecorder.status,
					Status:     http.StatusText(respRecorder.status),
					Method:     req.Method,
					Host:       req.Host,
					Path:       req.URL.Path,
					IP:         req.RemoteAddr,
					Headers:    "", // No headers available
					Time:       time.Now().UTC(),
				})
			}
		}
	})
}

// statusRecorder wraps a http.ResponseWriter to capture the status code
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}
