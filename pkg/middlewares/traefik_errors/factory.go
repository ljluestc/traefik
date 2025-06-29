package traefik_errors

import (
	"context"
	"net/http"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/middlewares"
)

// Factory is the middleware factory for TraefikErrors middleware.
type Factory struct {
	nameGenerator middlewares.NameGenerator
}

// NewFactory creates a new middleware factory.
func NewFactory(nameGenerator middlewares.NameGenerator) *Factory {
	return &Factory{nameGenerator: nameGenerator}
}

// New creates a new TraefikErrors middleware.
func (f *Factory) New(ctx context.Context, next http.Handler, config dynamic.TraefikErrors, name string) (http.Handler, error) {
	middlewareName := middlewares.GetQualifiedName(f.nameGenerator, name)
	return New(ctx, next, config, middlewareName)
}

// Factory is the middleware factory for TraefikErrors middleware.
type Factory struct {
	nameGenerator middlewares.NameGenerator
}

// NewFactory creates a new middleware factory.
func NewFactory(nameGenerator middlewares.NameGenerator) *Factory {
	return &Factory{nameGenerator: nameGenerator}
}

// New creates a new TraefikErrors middleware.
func (f *Factory) New(ctx context.Context, next http.Handler, config dynamic.TraefikErrors, name string) (http.Handler, error) {
	middlewareName := middlewares.GetQualifiedName(f.nameGenerator, name)
	return New(ctx, next, config, middlewareName)
}
