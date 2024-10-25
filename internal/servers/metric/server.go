package metric

import (
	"context"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	"net"
	"net/http"
	"strconv"
	"time"
)

// MetricServer represents a metric servers
type MetricServer struct {
	engine *gin.Engine
	http   http.Server
	config config.MetricServer
}

// NewMetricServer creates a new metric servers
func NewMetricServer(config config.MetricServer) *MetricServer {
	// Create a new metric servers
	i := &MetricServer{config: config}

	// Create a new gin engine
	if config.Certificate.Status {
		i.createHttpsServer()
	} else {
		i.createHttpServer()
	}
	return i
}

// createHttpServer creates a new HTTP servers
func (s *MetricServer) createHttpServer() {
	// Create a new HTTP servers
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	s.http = http.Server{
		Addr:         address,
		Handler:      s.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// createHttpServer creates a new HTTPS servers
func (s *MetricServer) createHttpsServer() {
	// Load the TLS certificate
	cert, err := tls.LoadX509KeyPair(
		s.config.Certificate.CertFile,
		s.config.Certificate.KeyFile,
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load TLS configuration")
	}

	// Create a new HTTPS servers
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	s.http = http.Server{
		Addr:         address,
		Handler:      s.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
		},
	}
}

// ListenAndServe starts the metric servers
func (s *MetricServer) ListenAndServe() error {
	log.Info().
		Str("host", s.config.Host).
		Int("port", s.config.Port).
		Bool("tls", s.config.Certificate.Status).
		Msg("Starting the rest servers")
	if s.config.Certificate.Status {
		return s.http.ListenAndServeTLS("", "")
	} else {
		return s.http.ListenAndServe()
	}
}

// Shutdown stops the metric servers
func (s *MetricServer) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down the metric servers")
	return s.http.Shutdown(ctx)
}
