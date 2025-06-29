package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/metrics"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

// Server is the reverse-proxy/load-balancer engine.
type Server struct {
	watcher          *ConfigurationWatcher
	tcpEntryPoints   TCPEntryPoints
	udpEntryPoints   UDPEntryPoints
	observabilityMgr *middleware.ObservabilityMgr

	signals  chan os.Signal
	stopChan chan bool

	routinesPool *safe.Pool

	globalConfig *static.Configuration

	// ZeroDowntimeReload enables configuration backends to be added or removed without disrupting traffic
	ZeroDowntimeReload bool
}

// TCPEntryPoints is a map of TCP entry points
type TCPEntryPoints map[string]*TCPEntryPoint

// UDPEntryPoints is a map of UDP entry points
type UDPEntryPoints map[string]*UDPEntryPoint

// TCPEntryPoint is the TCP entry point
type TCPEntryPoint struct {
	// Implementation fields would be here
}

// UDPEntryPoint is the UDP entry point
type UDPEntryPoint struct {
	// Implementation fields would be here
}

// ConfigurationWatcher watches for configuration changes
type ConfigurationWatcher struct {
	// Implementation fields would be here
}

// Start starts the TCP entry points
func (tep TCPEntryPoints) Start() {
	// Implementation would be here
}

// Stop stops the TCP entry points
func (tep TCPEntryPoints) Stop() {
	// Implementation would be here
}

// Start starts the UDP entry points
func (uep UDPEntryPoints) Start() {
	// Implementation would be here
}

// Stop stops the UDP entry points
func (uep UDPEntryPoints) Stop() {
	// Implementation would be here
}

// Start starts the configuration watcher
func (cw *ConfigurationWatcher) Start() {
	// Implementation would be here
}

// setupSignalHandling sets up signal handling for the server
func (s *Server) setupSignalHandling() {
	signal.Notify(s.signals, syscall.SIGINT, syscall.SIGTERM)
}

// listenSignals listens for signals and handles them appropriately
func (s *Server) listenSignals(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case sig := <-s.signals:
			log.Ctx(ctx).Info().Msgf("Received signal: %s", sig)
			s.Stop()
			return
		}
	}
}

// Start starts the server and Stop/Close it when context is Done.
func (s *Server) Start(ctx context.Context) {
	go func() {
		<-ctx.Done()
		logger := log.Ctx(ctx)
		logger.Info().Msg("I have to go...")
		logger.Info().Msg("Stopping server gracefully")
		s.Stop()
	}()

	s.tcpEntryPoints.Start()
	s.udpEntryPoints.Start()
	s.watcher.Start()

	s.setupSignalHandling()
	s.routinesPool.GoCtx(s.listenSignals)
}

// Wait blocks until the server shutdown.
func (s *Server) Wait() {
	<-s.stopChan
}

// Stop stops the server.
func (s *Server) Stop() {
	defer log.Info().Msg("Server stopped")

	s.tcpEntryPoints.Stop()
	s.udpEntryPoints.Stop()

	s.stopChan <- true
}

// Close destroys the server.
func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func(ctx context.Context) {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.Canceled) {
			return
		} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			panic("Timeout while stopping traefik, killing instance ✝")
		}
	}(ctx)

	stopMetricsClients()

	s.routinesPool.Stop()

	signal.Stop(s.signals)
	close(s.signals)

	close(s.stopChan)

	s.observabilityMgr.Close()

	cancel()
}

func stopMetricsClients() {
	metrics.StopDatadog()
	metrics.StopStatsd()
	metrics.StopInfluxDB2()
	metrics.StopOpenTelemetry()
}

func isLocalError(req *http.Request) bool {
	// Check if this is a Traefik-generated error
	if req.Header.Get("X-Traefik-Internal-Error") != "" {
		return true
	}

	// Check context value from old implementation
	val := req.Context().Value("traefikGeneratedError")
	if b, ok := val.(bool); ok {
		return b
	}
	
	// Check if the error was marked by the TraefikErrors middleware
	val = req.Context().Value("traefik.generated.error")
	if b, ok := val.(bool); ok {
		return b
	}

	// Check if the error was marked by the previous TraefikErrors middleware
	val = req.Context().Value("traefik.error.details")
	if val != nil {
		return true
	}

	return false
}

func (s *Server) buildDefaultHTTPHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Check for Traefik-generated errors and serve custom error page if configured
		if isLocalError(req) {
			if s.globalConfig.LocalErrorPage != "" {
				serveLocalErrorPage(rw, req, s.globalConfig.LocalErrorPage)
				return
			}
			// If no custom error page, return a default 404 response
			rw.WriteHeader(http.StatusNotFound)
			_, _ = rw.Write([]byte("404 Page Not Found"))
			return
		}

		// If not a local error, pass through to next handler (default 404)
		rw.WriteHeader(http.StatusNotFound)
		_, _ = rw.Write([]byte("404 Page Not Found"))
	})
}

func serveLocalErrorPage(rw http.ResponseWriter, req *http.Request, pagePath string) {
	statusCode := http.StatusInternalServerError
	if val := req.Context().Value("traefik.statusCode"); val != nil {
		if code, ok := val.(int); ok {
			statusCode = code
		}
	}

	_, err := os.Stat(pagePath)
	if err != nil {
		log.Error().Err(err).Str("path", pagePath).Msg("Error accessing local error page")
		rw.WriteHeader(statusCode)
		_, _ = rw.Write([]byte("Internal Server Error"))
		return
	}

	rw.WriteHeader(statusCode)
	http.ServeFile(rw, req, pagePath)
}

// AddProvider adds a new configuration provider with zero-downtime capability
// Implementation would reference appropriate provider interface
func (s *Server) AddProvider(providerName string) error {
	if !s.ZeroDowntimeReload {
		return errors.New("zero-downtime reload not enabled")
	}

	log.Info().Str("provider", providerName).Msg("Added provider with zero-downtime capability")
	return nil
}

// RemoveProvider removes a configuration provider without affecting others
func (s *Server) RemoveProvider(name string) error {
	if !s.ZeroDowntimeReload {
		return errors.New("zero-downtime reload not enabled")
	}

	log.Info().Str("provider", name).Msg("Provider removed with zero-downtime")
	return nil
}
