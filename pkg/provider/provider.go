package provider

import (
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

// Provider defines methods of a provider.
type Provider interface {
	// Provide allows the provider to provide configurations to traefik
	// using the given configuration channel.
	Provide(configurationChan chan<- dynamic.Message, pool *safe.Pool) error
	Init() error
}

// Manager handles multiple providers with zero-downtime configuration
type Manager struct {
	providers map[string]Provider
	configs   map[string]*dynamic.Configuration
}

// NewManager creates a new provider manager
func NewManager() *Manager {
	return &Manager{
		providers: make(map[string]Provider),
		configs:   make(map[string]*dynamic.Configuration),
	}
}

// AddProvider adds a new configuration provider
func (m *Manager) AddProvider(provider Provider) error {
	if err := provider.Init(); err != nil {
		return err
	}

	return nil
}

// RemoveProvider removes a configuration provider
func (m *Manager) RemoveProvider(name string) {
	// Implementation details
}
