package service

import (
	"errors"
	"net/http"
)

// HandlerFactory builds HTTP handlers.
type HandlerFactory struct {
	loadBalancer *ServiceLoadBalancer
	weighted     *ServiceWeightedRoundRobin
	mirroring    *ServiceMirroring
	failover     *ServiceFailover
}

// ServiceLoadBalancer is an interface for load balancers.
type ServiceLoadBalancer interface {
	Build() (http.Handler, error)
}

// ServiceWeightedRoundRobin is an interface for weighted round robin load balancers.
type ServiceWeightedRoundRobin interface {
	Build() (http.Handler, error)
}

// ServiceMirroring is an interface for mirroring services.
type ServiceMirroring interface {
	Build() (http.Handler, error)
}

// ServiceFailover is an interface for failover services.
type ServiceFailover interface {
	Build() (http.Handler, error)
}

// NewHandlerFactory creates a new HandlerFactory.
func NewHandlerFactory(loadBalancer *ServiceLoadBalancer, weighted *ServiceWeightedRoundRobin, mirroring *ServiceMirroring, failover *ServiceFailover) *HandlerFactory {
	return &HandlerFactory{
		loadBalancer: loadBalancer,
		weighted:     weighted,
		mirroring:    mirroring,
		failover:     failover,
	}
}

// Build builds a HTTP handler.
func (s *HandlerFactory) Build() (http.Handler, error) {
	if s.loadBalancer != nil {
		loadBalancer, err := (*s.loadBalancer).Build()
		if err != nil {
			return nil, err
		}

		return loadBalancer, nil
	}

	if s.weighted != nil {
		loadBalancer, err := (*s.weighted).Build()
		if err != nil {
			return nil, err
		}

		return loadBalancer, nil
	}

	if s.mirroring != nil {
		mirroring, err := (*s.mirroring).Build()
		if err != nil {
			return nil, err
		}

		return mirroring, nil
	}

	if s.failover != nil {
		failover, err := (*s.failover).Build()
		if err != nil {
			return nil, err
		}

		return failover, nil
	}

	return nil, errors.New("no service configuration found")
}

// HandleNotFoundError handles 404 errors as Traefik-generated errors to allow custom error pages
func HandleNotFoundError(req *http.Request, statusCode int) *http.Request {
	// Handle 404 errors as Traefik-generated errors to allow custom error pages
	if statusCode == http.StatusNotFound {
		req = markAsTraefikGeneratedError(req)
	}
	return req
}

// markAsTraefikGeneratedError marks a request as a Traefik-generated error
func markAsTraefikGeneratedError(req *http.Request) *http.Request {
	// Implementation would go here
	return req
}
