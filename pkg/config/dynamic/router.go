package dynamic

import (
	"github.com/traefik/traefik/v3/pkg/tls"
)

// Router holds the router configuration.
type Router struct {
	// EntryPoints defines the entry points by name.
	EntryPoints []string `json:"entryPoints,omitempty" toml:"entryPoints,omitempty" yaml:"entryPoints,omitempty" export:"true"`

	// Middlewares defines the list of middleware names.
	Middlewares []string `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`

	// Service defines the service name to use.
	Service string `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`

	// Rule defines the router's rule.
	Rule string `json:"rule,omitempty" toml:"rule,omitempty" yaml:"rule,omitempty" export:"true"`

	// RuleSyntax defines the rule syntax used.
	RuleSyntax string `json:"ruleSyntax,omitempty" toml:"ruleSyntax,omitempty" yaml:"ruleSyntax,omitempty" export:"true"`

	// Priority defines the router's priority.
	Priority int `json:"priority,omitempty" toml:"priority,omitempty" yaml:"priority,omitempty" export:"true"`

	// TLS defines the TLS configuration.
	TLS *RouterTLSConfig `json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`

	// Observability defines the observability configuration.
	Observability *RouterObservabilityConfig `json:"observability,omitempty" toml:"observability,omitempty" yaml:"observability,omitempty" export:"true"`
}

// RouterObservabilityConfig holds the router's observability configuration.
type RouterObservabilityConfig struct {
	// AccessLogs defines if access logs are enabled.
	AccessLogs *bool `json:"accessLogs,omitempty" toml:"accessLogs,omitempty" yaml:"accessLogs,omitempty" export:"true"`

	// Tracing defines if tracing is enabled.
	Tracing *bool `json:"tracing,omitempty" toml:"tracing,omitempty" yaml:"tracing,omitempty" export:"true"`

	// Metrics defines if metrics are enabled.
	Metrics *bool `json:"metrics,omitempty" toml:"metrics,omitempty" yaml:"metrics,omitempty" export:"true"`
}

// RouterTLSConfig holds the router's TLS configuration.
type RouterTLSConfig struct {
	// Options defines the reference to TLS options.
	Options string `json:"options,omitempty" toml:"options,omitempty" yaml:"options,omitempty" export:"true"`

	// CertResolver defines the certificate resolver name.
	CertResolver string `json:"certResolver,omitempty" toml:"certResolver,omitempty" yaml:"certResolver,omitempty" export:"true"`

	// Domains defines the list of domains that will be used for TLS certificate generation.
	Domains []tls.Domain `json:"domains,omitempty" toml:"domains,omitempty" yaml:"domains,omitempty" export:"true"`
}

// Model holds the model configuration.
type Model struct {
	// Middlewares defines the list of middleware names.
	Middlewares []string `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`

	// TLS defines the TLS configuration.
	TLS *RouterTLSConfig `json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`

	// Observability defines the observability configuration.
	Observability RouterObservabilityConfig `json:"observability,omitempty" toml:"observability,omitempty" yaml:"observability,omitempty" export:"true"`
}
