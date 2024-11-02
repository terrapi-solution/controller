package grpc

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/controller/internal/core"
	services "github.com/terrapi-solution/controller/router/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

// GrpcServer represents a gRPC servers
type GrpcServer struct {
	services *core.Core
	config   config.GrpcServer
	server   *grpc.Server
	listener *net.Listener
}

// NewGRPCServer initializes and returns a gRPC servers
func NewGRPCServer(cfg *config.Config, coreService *core.Core) *GrpcServer {
	// Create a new gRPC servers
	i := &GrpcServer{
		services: coreService,
		config:   cfg.Servers.Grpc,
	}

	i.server = i.createGrpcServer()

	// Register all gRPC services
	services.NewGrpcService(i.server)

	// Enable gRPC servers reflection
	reflection.Register(i.server)

	// Return the gRPC servers
	return i
}

// createListener creates a new network listener
func (s *GrpcServer) createListener() {
	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	s.listener = &lis
}

// NewGrpcServer creates a new grpc server
func (s *GrpcServer) createGrpcServer() *grpc.Server {
	if s.config.Certificate.Status {
		tlsConfig, err := s.getTransportCredentials(
			s.config.Certificate.CertFile,
			s.config.Certificate.KeyFile,
			s.config.Certificate.CaFile)
		if err != nil {
			log.Panic().Err(err).Msg("failed to load TLS configuration")
		}
		return grpc.NewServer(grpc.Creds(tlsConfig))
	} else {
		return grpc.NewServer()
	}
}

// ListenAndServe starts the grpc server
func (s *GrpcServer) ListenAndServe() error {
	log.Info().
		Str("host", s.config.Host).
		Int("port", s.config.Port).
		Bool("tls", s.config.Certificate.Status).
		Msg("Starting the grpc servers")
	s.createListener()
	return s.server.Serve(*s.listener)
}

// Shutdown stops the grpc server
func (s *GrpcServer) Shutdown() {
	log.Info().Msg("Shutting down the grpc servers")
	s.server.GracefulStop()
}
