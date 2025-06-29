package dynamic

import (
	"github.com/traefik/traefik/v3/pkg/tls"
)

// +k8s:deepcopy-gen=true

// HealthErrorLog holds configuration for extended error logging in the /health API.
type HealthErrorLog struct {
	IncludeHeaders bool `json:"includeHeaders,omitempty" toml:"includeHeaders,omitempty" yaml:"includeHeaders,omitempty"`
	IncludeIP      bool `json:"includeIP,omitempty" toml:"includeIP,omitempty" yaml:"includeIP,omitempty"`
}

// Configuration is the root of the dynamic configuration.
type Configuration struct {
	HTTP           *HTTPConfiguration `json:"http,omitempty" toml:"http,omitempty" yaml:"http,omitempty" export:"true"`
	TCP            *TCPConfiguration  `json:"tcp,omitempty" toml:"tcp,omitempty" yaml:"tcp,omitempty" export:"true"`
	UDP            *UDPConfiguration  `json:"udp,omitempty" toml:"udp,omitempty" yaml:"udp,omitempty" export:"true"`
	TLS            *TLSConfiguration  `json:"tls,omitempty" toml:"tls,omitempty" yaml:"tls,omitempty" export:"true"`
	HealthErrorLog *HealthErrorLog    `json:"healthErrorLog,omitempty" toml:"healthErrorLog,omitempty" yaml:"healthErrorLog,omitempty" export:"true"`
}

// HTTPConfiguration contains all the HTTP configuration parts.
type HTTPConfiguration struct {
	Routers     map[string]*Router     `json:"routers,omitempty" toml:"routers,omitempty" yaml:"routers,omitempty" export:"true"`
	Services    map[string]*Service    `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
	Middlewares map[string]*Middleware `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`
}

// TCPConfiguration contains all the TCP configuration parts.
type TCPConfiguration struct {
	Routers  map[string]*TCPRouter  `json:"routers,omitempty" toml:"routers,omitempty" yaml:"routers,omitempty" export:"true"`
	Services map[string]*TCPService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
}

// UDPConfiguration contains all the UDP configuration parts.
type UDPConfiguration struct {
	Routers  map[string]*UDPRouter  `json:"routers,omitempty" toml:"routers,omitempty" yaml:"routers,omitempty" export:"true"`
	Services map[string]*UDPService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
}

// HTTPMiddleware holds the middleware configuration.
type HTTPMiddleware struct {
	// Configuration placeholder - will be implemented with specific middleware types
}

// TCPRouter holds the TCP router configuration.
type TCPRouter struct {
	// Configuration placeholder
}

// TCPService holds the TCP service configuration.
type TCPService struct {
	LoadBalancer *TCPServersLoadBalancer `json:"loadBalancer,omitempty" toml:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty" export:"true"`
	Weighted     *TCPWeightedRoundRobin  `json:"weighted,omitempty" toml:"weighted,omitempty" yaml:"weighted,omitempty" export:"true"`
}

// UDPRouter holds the UDP router configuration.
type UDPRouter struct {
	// Configuration placeholder
}

// UDPService holds the UDP service configuration.
type UDPService struct {
	LoadBalancer *UDPServersLoadBalancer `json:"loadBalancer,omitempty" toml:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty" export:"true"`
	Weighted     *UDPWeightedRoundRobin  `json:"weighted,omitempty" toml:"weighted,omitempty" yaml:"weighted,omitempty" export:"true"`
}

// +k8s:deepcopy-gen=true

// TLSConfiguration contains all the configuration parameters of a TLS connection.
type TLSConfiguration struct {
	Certificates []*tls.CertAndStores   `json:"certificates,omitempty"  toml:"certificates,omitempty" yaml:"certificates,omitempty" label:"-" export:"true"`
	Options      map[string]tls.Options `json:"options,omitempty" toml:"options,omitempty" yaml:"options,omitempty" label:"-" export:"true"`
	Stores       map[string]tls.Store   `json:"stores,omitempty" toml:"stores,omitempty" yaml:"stores,omitempty" export:"true"`
}

// +k8s:deepcopy-gen=true

// Message holds configuration information exchanged between parts of traefik.
type Message struct {
	ProviderName  string
	Configuration *Configuration
}

// +k8s:deepcopy-gen=true

// Configurations is for currentConfigurations Map.
type Configurations map[string]*Configuration
