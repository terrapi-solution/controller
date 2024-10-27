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
func (r *RestServer) createHttpServer() {
	// Create a new HTTP servers
	address := net.JoinHostPort(r.config.Host, strconv.Itoa(r.config.Port))
	r.http = http.Server{
		Addr:         address,
		Handler:      r.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// createHttpServer creates a new HTTPS server
func (r *RestServer) createHttpsServer() {
	// Load the TLS certificate
	cert, err := tls.LoadX509KeyPair(
		r.config.Certificate.CertFile,
		r.config.Certificate.KeyFile,
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load TLS configuration")
	}

	// Create a new HTTPS server
	address := net.JoinHostPort(r.config.Host, strconv.Itoa(r.config.Port))
	r.http = http.Server{
		Addr:         address,
		Handler:      r.loadRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: r.curves(),
			CipherSuites:     r.ciphers(),
			Certificates:     []tls.Certificate{cert},
		},
	}
}

// ListenAndServe starts the HTTP server
func (r *RestServer) ListenAndServe() error {
	log.Info().
		Str("host", r.config.Host).
		Int("port", r.config.Port).
		Bool("tls", r.config.Certificate.Status).
		Msg("Starting the rest servers")
	if r.config.Certificate.Status {
		return r.http.ListenAndServeTLS("", "")
	} else {
		return r.http.ListenAndServe()
	}
}

// Shutdown stops the rest server
func (r *RestServer) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down the rest servers")
	return r.http.Shutdown(ctx)
}
