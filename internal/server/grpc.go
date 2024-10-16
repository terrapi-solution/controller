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
	"k8s.io/utils/env"
	"os"

	"google.golang.org/grpc/reflection"
)

var MaxGRPCMessageSize int

func init() {
	var err error
	MaxGRPCMessageSize, err = env.GetInt("GRPC_MESSAGE_SIZE", 100*1024*1024)
	if err != nil {
		log.Fatal().Err(err).
			Msg("GRPC_MESSAGE_SIZE environment variable must be set as an integer")
	}
}

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
	if cfg.Server.Certificates.Status {
		tlsConfig, err := loadTLSConfig(
			cfg.Server.Certificates.CertFile,
			cfg.Server.Certificates.KeyFile,
			cfg.Server.Certificates.CaFile)
		if err != nil {
			log.Panic().Err(err).Msg("failed to load TLS configuration")
		}
		return grpc.NewServer(grpc.Creds(tlsConfig))
	} else {
		return grpc.NewServer()
	}
}

// loadTlSConfig loads the TLS configuration
func loadTLSConfig(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load server certification: %w", err)
	}

	data, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("faild to read CA certificate: %w", err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(data) {
		return nil, fmt.Errorf("unable to append the CA certificate to CA pool")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    pool,
	}
	return credentials.NewTLS(tlsConfig), nil
}
