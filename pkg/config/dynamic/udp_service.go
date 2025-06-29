package dynamic

// UDPServersLoadBalancer holds the UDP servers load balancer configuration.
type UDPServersLoadBalancer struct {
	// Servers defines the list of UDP server URLs.
	Servers []UDPServer `json:"servers,omitempty" toml:"servers,omitempty" yaml:"servers,omitempty" export:"true"`
}

// UDPServer holds a UDP server configuration.
type UDPServer struct {
	// Address defines the UDP server address.
	Address string `json:"address,omitempty" toml:"address,omitempty" yaml:"address,omitempty" export:"true"`
}

// UDPWeightedRoundRobin holds the UDP weighted round robin configuration.
type UDPWeightedRoundRobin struct {
	// Services defines the list of UDP services.
	Services []UDPWRRService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
}

// UDPWRRService holds a UDP WRR service configuration.
type UDPWRRService struct {
	// Name defines the service name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	
	// Weight defines the service weight.
	Weight *int `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" export:"true"`
}
