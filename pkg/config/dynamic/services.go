package dynamic

import (
	"crypto/tls"
)

// ServersTransport holds the server transport configuration.
type ServersTransport struct {
	ServerName          string              `json:"serverName,omitempty" toml:"serverName,omitempty" yaml:"serverName,omitempty" export:"true"`
	InsecureSkipVerify  bool                `json:"insecureSkipVerify,omitempty" toml:"insecureSkipVerify,omitempty" yaml:"insecureSkipVerify,omitempty" export:"true"`
	RootCAs             []string            `json:"rootCAs,omitempty" toml:"rootCAs,omitempty" yaml:"rootCAs,omitempty" export:"true"`
	Certificates        []tls.Certificate   `json:"certificates,omitempty" toml:"certificates,omitempty" yaml:"certificates,omitempty" export:"true"`
	MaxIdleConnsPerHost int                 `json:"maxIdleConnsPerHost,omitempty" toml:"maxIdleConnsPerHost,omitempty" yaml:"maxIdleConnsPerHost,omitempty" export:"true"`
	ForwardingTimeouts  *ForwardingTimeouts `json:"forwardingTimeouts,omitempty" toml:"forwardingTimeouts,omitempty" yaml:"forwardingTimeouts,omitempty" export:"true"`
	DisableHTTP2        bool                `json:"disableHTTP2,omitempty" toml:"disableHTTP2,omitempty" yaml:"disableHTTP2,omitempty" export:"true"`
	PeerCertURI         string              `json:"peerCertURI,omitempty" toml:"peerCertURI,omitempty" yaml:"peerCertURI,omitempty" export:"true"`
	Spiffe              *Spiffe             `json:"spiffe,omitempty" toml:"spiffe,omitempty" yaml:"spiffe,omitempty" export:"true"`
}

// Service holds the service configuration.
type Service struct {
	LoadBalancer *ServersLoadBalancer `json:"loadBalancer,omitempty" toml:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty" export:"true"`
	Weighted     *WeightedRoundRobin  `json:"weighted,omitempty" toml:"weighted,omitempty" yaml:"weighted,omitempty" export:"true"`
	Mirroring    *Mirroring           `json:"mirroring,omitempty" toml:"mirroring,omitempty" yaml:"mirroring,omitempty" export:"true"`
	Failover     *Failover            `json:"failover,omitempty" toml:"failover,omitempty" yaml:"failover,omitempty" export:"true"`
}

// Spiffe holds the SPIFFE configuration.
type Spiffe struct {
	IDs         []string `json:"ids,omitempty" toml:"ids,omitempty" yaml:"ids,omitempty" export:"true"`
	TrustDomain string   `json:"trustDomain,omitempty" toml:"trustDomain,omitempty" yaml:"trustDomain,omitempty" export:"true"`
}

// Sticky holds the sticky session configuration.
type Sticky struct {
	Cookie *Cookie `json:"cookie,omitempty" toml:"cookie,omitempty" yaml:"cookie,omitempty" export:"true"`
}

// WRRService holds the weighted round robin service configuration.
type WRRService struct {
	Name   string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	Weight *int   `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" export:"true"`
}

// WeightedRoundRobin holds the weighted round robin configuration.
type WeightedRoundRobin struct {
	Services    []WRRService `json:"services,omitempty" toml:"services,omitempty" yaml:"services,omitempty" export:"true"`
	Sticky      *Sticky      `json:"sticky,omitempty" toml:"sticky,omitempty" yaml:"sticky,omitempty" export:"true"`
	HealthCheck *HealthCheck `json:"healthCheck,omitempty" toml:"healthCheck,omitempty" yaml:"healthCheck,omitempty" export:"true"`
}

// ForwardingTimeouts holds the forwarding timeouts configuration.
type ForwardingTimeouts struct {
	DialTimeout           string `json:"dialTimeout,omitempty" toml:"dialTimeout,omitempty" yaml:"dialTimeout,omitempty" export:"true"`
	ResponseHeaderTimeout string `json:"responseHeaderTimeout,omitempty" toml:"responseHeaderTimeout,omitempty" yaml:"responseHeaderTimeout,omitempty" export:"true"`
	IdleConnTimeout       string `json:"idleConnTimeout,omitempty" toml:"idleConnTimeout,omitempty" yaml:"idleConnTimeout,omitempty" export:"true"`
}

// ServersLoadBalancer holds the server load balancer configuration.
type ServersLoadBalancer struct {
	Servers            []Server            `json:"servers,omitempty" toml:"servers,omitempty" yaml:"servers,omitempty" export:"true"`
	HealthCheck        *HealthCheck        `json:"healthCheck,omitempty" toml:"healthCheck,omitempty" yaml:"healthCheck,omitempty" export:"true"`
	Sticky             *Sticky             `json:"sticky,omitempty" toml:"sticky,omitempty" yaml:"sticky,omitempty" export:"true"`
	Strategy           string              `json:"strategy,omitempty" toml:"strategy,omitempty" yaml:"strategy,omitempty" export:"true"`
	PassHostHeader     *bool               `json:"passHostHeader,omitempty" toml:"passHostHeader,omitempty" yaml:"passHostHeader,omitempty" export:"true"`
	ResponseForwarding *ResponseForwarding `json:"responseForwarding,omitempty" toml:"responseForwarding,omitempty" yaml:"responseForwarding,omitempty" export:"true"`
	ServersTransport   string              `json:"serversTransport,omitempty" toml:"serversTransport,omitempty" yaml:"serversTransport,omitempty" export:"true"`
}

// Server holds the server configuration.
type Server struct {
	URL          string `json:"url,omitempty" toml:"url,omitempty" yaml:"url,omitempty" export:"true"`
	Weight       *int   `json:"weight,omitempty" toml:"weight,omitempty" yaml:"weight,omitempty" export:"true"`
	PreservePath bool   `json:"preservePath,omitempty" toml:"preservePath,omitempty" yaml:"preservePath,omitempty" export:"true"`
	Fenced       bool   `json:"fenced,omitempty" toml:"fenced,omitempty" yaml:"fenced,omitempty" export:"true"`
}

// HealthCheck holds the health check configuration.
type HealthCheck struct {
	Scheme          string            `json:"scheme,omitempty" toml:"scheme,omitempty" yaml:"scheme,omitempty" export:"true"`
	Path            string            `json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty" export:"true"`
	Port            int               `json:"port,omitempty" toml:"port,omitempty" yaml:"port,omitempty" export:"true"`
	Interval        string            `json:"interval,omitempty" toml:"interval,omitempty" yaml:"interval,omitempty" export:"true"`
	Timeout         string            `json:"timeout,omitempty" toml:"timeout,omitempty" yaml:"timeout,omitempty" export:"true"`
	Headers         map[string]string `json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty" export:"true"`
	FollowRedirects *bool             `json:"followRedirects,omitempty" toml:"followRedirects,omitempty" yaml:"followRedirects,omitempty" export:"true"`
	Mode            string            `json:"mode,omitempty" toml:"mode,omitempty" yaml:"mode,omitempty" export:"true"`
	Method          string            `json:"method,omitempty" toml:"method,omitempty" yaml:"method,omitempty" export:"true"`
	Status          int               `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`
	Hostname        string            `json:"hostname,omitempty" toml:"hostname,omitempty" yaml:"hostname,omitempty" export:"true"`
}

// Mirroring holds the mirroring configuration.
type Mirroring struct {
	Service     string          `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	MaxBodySize *int64          `json:"maxBodySize,omitempty" toml:"maxBodySize,omitempty" yaml:"maxBodySize,omitempty" export:"true"`
	Mirrors     []MirrorService `json:"mirrors,omitempty" toml:"mirrors,omitempty" yaml:"mirrors,omitempty" export:"true"`
}

// MirrorService holds the mirrored service configuration.
type MirrorService struct {
	Name    string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	Percent int    `json:"percent,omitempty" toml:"percent,omitempty" yaml:"percent,omitempty" export:"true"`
}

// Failover holds the failover configuration.
type Failover struct {
	Service     string       `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`
	Fallback    string       `json:"fallback,omitempty" toml:"fallback,omitempty" yaml:"fallback,omitempty" export:"true"`
	HealthCheck *HealthCheck `json:"healthCheck,omitempty" toml:"healthCheck,omitempty" yaml:"healthCheck,omitempty" export:"true"`
}

// Cookie holds the cookie configuration.
type Cookie struct {
	Name     string `json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" export:"true"`
	Secure   bool   `json:"secure,omitempty" toml:"secure,omitempty" yaml:"secure,omitempty" export:"true"`
	HTTPOnly bool   `json:"httpOnly,omitempty" toml:"httpOnly,omitempty" yaml:"httpOnly,omitempty" export:"true"`
	SameSite string `json:"sameSite,omitempty" toml:"sameSite,omitempty" yaml:"sameSite,omitempty" export:"true"`
}

// ResponseForwarding holds the response forwarding configuration.
type ResponseForwarding struct {
	FlushInterval string `json:"flushInterval,omitempty" toml:"flushInterval,omitempty" yaml:"flushInterval,omitempty" export:"true"`
}
