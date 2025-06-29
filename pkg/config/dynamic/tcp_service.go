package dynamic

import (
	"github.com/traefik/paerser/types"
)

// TCPServersLoadBalancer holds the TCP servers load balancer configuration.
type TCPServersLoadBalancer struct {
	// Servers defines the list of TCP server URLs.
	Servers []TCPServer `json:"servers,omitempty" toml:"servers,omitempty" yaml:"servers,omitempty" export:"true"`
	
	// TerminationDelay defines the delay to wait before closing connections.
	TerminationDelay *int `json:"terminationDelay,omitempty" toml:"terminationDelay,omitempty" yaml:"terminationDelay,omitempty" export:"true"`
	
	// ProxyProtocol defines the ProxyProtocol configuration.
	ProxyProtocol *ProxyProtocol `json:"proxyProtocol,omitempty" toml:"proxyProtocol,omitempty" yaml:"proxyProtocol,omitempty" export:"true"`
}

// TCPServer holds a TCP server configuration.
type TCPServer struct {
	// Address defines the TCP server address.
	Address string `json:"address,omitempty" toml:"address,omitempty" yaml:"address,omitempty" export:"true"`
}

// ProxyProtocol holds the ProxyProtocol configuration.
type ProxyProtocol struct {
	// Version defines the Protocol version.
	Version int `json:"version,omitempty" toml:"version,omitempty" yaml:"version,omitempty" export:"true"`
}

// TCPWeightedRoundRobin holds the TCP weighted round robin configuration.
type TCPWeightedRoundRobin struct {
	// Services defines the list of TCP services.
	Services []TCPWRRService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
}

// TCPWRRService holds a TCP WRR service configuration.
type TCPWRRService struct {
	// Name defines the service name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	
	// Weight defines the service weight.
	Weight *int `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" export:"true"`
}
