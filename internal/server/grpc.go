package server

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/controller"
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/controller/internal/service"
	"github.com/terrapi-solution/protocol/activity"
	"github.com/terrapi-solution/protocol/deployment"
	"github.com/terrapi-solution/protocol/health"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

// NewGRPCServer initializes and returns a gRPC server
func NewGRPCServer(cfg *config.Config) *grpc.Server {
	// Create a new gRPC server
	server := NewGrpcServer(cfg)

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
func NewGrpcServer(cfg *config.Config) *grpc.Server {
	if cfg.Server.Mode != "OIDC" {
		return grpc.NewServer()
	}

	// Initialise auth service & interceptor
	authSvc, err := service.NewAuthService(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize auth service")
	}
	interceptor, err := service.NewAuthInterceptorService(authSvc)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize interceptor")
	}

	// Create a new grpc server with the interceptor
	return grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryAuthMiddleware),
	)
}
