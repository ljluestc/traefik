package dynamic

// Errors holds the custom error pages configuration.
type Errors struct {
	Status  []string `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`
	Service string   `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	Query   string   `json:"query,omitempty" toml:"query,omitempty" yaml:"query,omitempty" export:"true"`
}

// TraefikErrors holds the custom error pages configuration for Traefik-generated errors.
type TraefikErrors struct {
	// Status defines the HTTP status codes to match (e.g., ["404", "500-599"])
	Status []string `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`
	
	// Service defines the name of the service that will serve the error pages
	Service string `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	
	// Query defines the URL path used to query the error page service
	// The {status} placeholder will be replaced by the actual error code
	Query string `json:"query,omitempty" toml:"query,omitempty" yaml:"query,omitempty" export:"true"`
}
