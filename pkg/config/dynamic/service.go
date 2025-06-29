package dynamic

// Service holds the HTTP service configuration.
type Service struct {
	// LoadBalancer holds the load balancer configuration.
	LoadBalancer *ServersLoadBalancer `json:"loadBalancer,omitempty" toml:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty" export:"true"`
	
	// Weighted holds the weighted round robin configuration.
	Weighted *WeightedRoundRobin `json:"weighted,omitempty" toml:"weighted,omitempty" yaml:"weighted,omitempty" export:"true"`
	
	// Mirroring holds the mirroring configuration.
	Mirroring *Mirroring `json:"mirroring,omitempty" toml:"mirroring,omitempty" yaml:"mirroring,omitempty" export:"true"`
	
	// Failover holds the failover configuration.
	Failover *Failover `json:"failover,omitempty" toml:"failover,omitempty" yaml:"failover,omitempty" export:"true"`
}

// ServersLoadBalancer holds the servers load balancer configuration.
type ServersLoadBalancer struct {
	// Servers defines the list of server URLs.
	Servers []Server `json:"servers,omitempty" toml:"servers,omitempty" yaml:"servers,omitempty" export:"true"`
	
	// PassHostHeader defines whether to pass the host header.
	PassHostHeader *bool `json:"passHostHeader,omitempty" toml:"passHostHeader,omitempty" yaml:"passHostHeader,omitempty" export:"true"`
	
	// ServersTransport defines the transport to use for the servers.
	ServersTransport string `json:"serversTransport,omitempty" toml:"serversTransport,omitempty" yaml:"serversTransport,omitempty" export:"true"`
	
	// Sticky defines the sticky sessions configuration.
	Sticky *Sticky `json:"sticky,omitempty" toml:"sticky,omitempty" yaml:"sticky,omitempty" export:"true"`
	
	// HealthCheck defines the health check configuration.
	HealthCheck *ServerHealthCheck `json:"healthCheck,omitempty" toml:"healthCheck,omitempty" yaml:"healthCheck,omitempty" export:"true"`
	
	// ResponseForwarding defines the response forwarding configuration.
	ResponseForwarding *ResponseForwarding `json:"responseForwarding,omitempty" toml:"responseForwarding,omitempty" yaml:"responseForwarding,omitempty" export:"true"`
}

// Server holds a server configuration.
type Server struct {
	// URL defines the server URL.
	URL string `json:"url,omitempty" toml:"url,omitempty" yaml:"url,omitempty" export:"true"`
}

// Sticky holds the sticky sessions configuration.
type Sticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *Cookie `json:"cookie,omitempty" toml:"cookie,omitempty" yaml:"cookie,omitempty" export:"true"`
}

// Cookie holds the sticky cookie configuration.
type Cookie struct {
	// Name defines the cookie name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	
	// Secure defines whether the cookie is secure.
	Secure bool `json:"secure,omitempty" toml:"secure,omitempty" yaml:"secure,omitempty" export:"true"`
	
	// HTTPOnly defines whether the cookie is HTTP only.
	HTTPOnly bool `json:"httpOnly,omitempty" toml:"httpOnly,omitempty" yaml:"httpOnly,omitempty" export:"true"`
}

// ResponseForwarding holds the response forwarding configuration.
type ResponseForwarding struct {
	// FlushInterval defines the interval in milliseconds between flushes to the client.
	FlushInterval string `json:"flushInterval,omitempty" toml:"flushInterval,omitempty" yaml:"flushInterval,omitempty" export:"true"`
}

// WeightedRoundRobin holds the weighted round robin configuration.
type WeightedRoundRobin struct {
	// Services defines the list of services.
	Services []WRRService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
	
	// Sticky defines the sticky sessions configuration.
	Sticky *Sticky `json:"sticky,omitempty" toml:"sticky,omitempty" yaml:"sticky,omitempty" export:"true"`
}

// WRRService holds a service configuration.
type WRRService struct {
	// Name defines the service name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	
	// Weight defines the service weight.
	Weight *int `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" export:"true"`
}

// Mirroring holds the mirroring configuration.
type Mirroring struct {
	// Service defines the main service name.
	Service string `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	
	// MaxBodySize defines the maximum body size in MB.
	MaxBodySize *int64 `json:"maxBodySize,omitempty" toml:"maxBodySize,omitempty" yaml:"maxBodySize,omitempty" export:"true"`
	
	// Mirrors defines the list of mirrors.
	Mirrors []Mirror `json:"mirrors,omitempty" toml:"mirrors,omitempty" yaml:"mirrors,omitempty" export:"true"`
}

// Mirror holds a mirror configuration.
type Mirror struct {
	// Name defines the mirror service name.
	Name string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	
	// Percent defines the percentage of requests to mirror.
	Percent int `json:"percent,omitempty" toml:"percent,omitempty" yaml:"percent,omitempty" export:"true"`
}

// Failover holds the failover configuration.
type Failover struct {
	// Service defines the fallback service name.
	Service string `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	
	// Fallback defines the fallback service name.
	Fallback string `json:"fallback,omitempty" toml:"fallback,omitempty" yaml:"fallback,omitempty" export:"true"`
}
