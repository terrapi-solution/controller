package rest

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

// RestServer represents a REST server
type RestServer struct {
	engine *gin.Engine
	http   http.Server
	config config.RestServer
}

// NewRestServer creates a new rest server
func NewRestServer(config config.RestServer) *RestServer {
	i := &RestServer{config: config}
	if config.Certificate.Status {
		i.createHttpsServer()
	} else {
		i.createHttpServer()
	}
	return i
}

// createHttpServer creates a new HTTP server
func (s *RestServer) createHttpServer() {
	// Create a new HTTP servers
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	s.http = http.Server{
		Addr:         address,
		Handler:      s.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// createHttpServer creates a new HTTPS server
func (s *RestServer) createHttpsServer() {
	// Load the TLS certificate
	cert, err := tls.LoadX509KeyPair(
		s.config.Certificate.CertFile,
		s.config.Certificate.KeyFile,
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load TLS configuration")
	}

	// Create a new HTTPS server
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	s.http = http.Server{
		Addr:         address,
		Handler:      s.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: s.curves(),
			CipherSuites:     s.ciphers(),
			Certificates:     []tls.Certificate{cert},
		},
	}
}

// ListenAndServe starts the HTTP server
func (s *RestServer) ListenAndServe() error {
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

// Shutdown stops the rest server
func (s *RestServer) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down the rest servers")
	return s.http.Shutdown(ctx)
}
