package dynamic

import (
	"github.com/traefik/paerser/types"
)

// ServerHealthCheck holds the health check configuration for a server.
type ServerHealthCheck struct {
	// Scheme defines the scheme used for the health check.
	Scheme string `json:"scheme,omitempty" toml:"scheme,omitempty" yaml:"scheme,omitempty" export:"true"`

	// Mode defines the health check mode.
	Mode string `json:"mode,omitempty" toml:"mode,omitempty" yaml:"mode,omitempty" export:"true"`

	// Path defines the URL path for the health check.
	Path string `json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty" export:"true"`

	// Method defines the HTTP method to use for the health check.
	Method string `json:"method,omitempty" toml:"method,omitempty" yaml:"method,omitempty" export:"true"`

	// Status defines the HTTP status code to expect for a successful health check.
	Status int `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`

	// Port defines the port to use for the health check.
	Port int `json:"port,omitempty" toml:"port,omitempty" yaml:"port,omitempty" export:"true"`

	// Interval defines the interval between health checks.
	Interval string `json:"interval,omitempty" toml:"interval,omitempty" yaml:"interval,omitempty" export:"true"`

	// UnhealthyInterval defines the interval for unhealthy targets.
	UnhealthyInterval *types.Duration `json:"unhealthyInterval,omitempty" toml:"unhealthyInterval,omitempty" yaml:"unhealthyInterval,omitempty" export:"true"`

	// Timeout defines the timeout for the health check request.
	Timeout string `json:"timeout,omitempty" toml:"timeout,omitempty" yaml:"timeout,omitempty" export:"true"`

	// Hostname defines the hostname to use for the health check.
	Hostname string `json:"hostname,omitempty" toml:"hostname,omitempty" yaml:"hostname,omitempty" export:"true"`

	// FollowRedirects defines whether to follow redirects during the health check.
	FollowRedirects *bool `json:"followRedirects,omitempty" toml:"followRedirects,omitempty" yaml:"followRedirects,omitempty" export:"true"`

	// Headers defines the HTTP headers to set in the health check request.
	Headers map[string]string `json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty" export:"true"`
}
