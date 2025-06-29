package dynamic

// BackendConfiguration manages the state of dynamic configuration backends
type BackendConfiguration struct {
	// Enabled determines if this backend is currently active
	Enabled bool `json:"enabled,omitempty" toml:"enabled,omitempty" yaml:"enabled,omitempty" export:"true"`
	
	// RefreshInterval is how often to check for new configuration
	RefreshInterval string `json:"refreshInterval,omitempty" toml:"refreshInterval,omitempty" yaml:"refreshInterval,omitempty" export:"true"`
	
	// LastRefresh contains the timestamp of the last successful refresh
	LastRefresh string `json:"lastRefresh,omitempty" toml:"lastRefresh,omitempty" yaml:"lastRefresh,omitempty" export:"true"`
	
	// LastError contains the last error encountered, if any
	LastError string `json:"lastError,omitempty" toml:"lastError,omitempty" yaml:"lastError,omitempty" export:"true"`
	
	// ZeroDowntimeReload enables independent backend handling
	ZeroDowntimeReload bool `json:"zeroDowntimeReload,omitempty" toml:"zeroDowntimeReload,omitempty" yaml:"zeroDowntimeReload,omitempty" export:"true"`
}

// BackendProvider defines a configuration provider
type BackendProvider struct {
	// Type specifies the provider type (consul, marathon, etc.)
	Type string `json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty" export:"true"`
	
	// Configuration contains the backend-specific configuration
	Configuration *BackendConfiguration `json:"configuration,omitempty" toml:"configuration,omitempty" yaml:"configuration,omitempty" export:"true"`
	
	// Status tracks the current state of this provider
	Status string `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`
}

// BackendManager handles multiple configuration backends
type BackendManager struct {
	// Providers contains all active configuration providers
	Providers map[string]*BackendProvider `json:"providers,omitempty" toml:"providers,omitempty" yaml:"providers,omitempty" export:"true"`
	
	// GlobalZeroDowntime enables zero-downtime reload globally
	GlobalZeroDowntime bool `json:"globalZeroDowntime,omitempty" toml:"globalZeroDowntime,omitempty" yaml:"globalZeroDowntime,omitempty" export:"true"`
}
