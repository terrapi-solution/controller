package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/controller"
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/protocol/activity"
	"github.com/terrapi-solution/protocol/deployment"
	"github.com/terrapi-solution/protocol/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"

	"google.golang.org/grpc/reflection"
)

// NewGRPCServer initializes and returns a gRPC server
func NewGRPCServer(cfg *config.Config) *grpc.Server {
	// Create a new gRPC server
	server := newGrpcServer(cfg)

	// Register all gRPC services
	deployment.RegisterDeploymentServiceServer(server, &controller.DeploymentServer{})
	activity.RegisterActivityServiceServer(server, &controller.ActivityServer{})
	health.RegisterHealthServiceServer(server, &controller.HealthServer{})

	// Enable gRPC server reflection
	reflection.Register(server)

	// Return the gRPC server
	return server
}

// NewGrpcServer creates a new grpc server
func newGrpcServer(cfg *config.Config) *grpc.Server {
	// Load the TLS certificate
	tlsConfig, err := loadTlSConfig(
		cfg.Server.Certificates.CaFile,
		cfg.Server.Certificates.KeyFile,
		cfg.Server.Certificates.CaFile)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load TLS configuration")
	}

	// Create a new grpc server with the interceptor
	return grpc.NewServer(grpc.Creds(tlsConfig))
}

// loadTlSConfig loads the TLS configuration
func loadTlSConfig(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load server certification: %w", err)
	}

	data, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("faild to read CA certificate: %w", err)
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(data) {
		return nil, fmt.Errorf("unable to append the CA certificate to CA pool")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    capool,
	}
	return credentials.NewTLS(tlsConfig), nil
}
