package traefik_errors

import (
	"time"
)

// ErrorDetails contains detailed information about errors for the health API
type ErrorDetails struct {
	StatusCode    int       `json:"status_code"`
	Status        string    `json:"status"`
	Method        string    `json:"method"`
	Host          string    `json:"host"`
	Path          string    `json:"path"`
	IP            string    `json:"ip"`
	Headers       string    `json:"headers"`
	Time          time.Time `json:"time"`
}
