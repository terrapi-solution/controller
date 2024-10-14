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

// NewGRPCServer creates a new grpc server
func NewGRPCServer(cfg *config.Config) *grpc.Server {
	// Create a new grpc server
	server := getGrpcServer(cfg)

	// Register the service with the server
	deployment.RegisterDeploymentServiceServer(server, &controller.DeploymentServer{})
	activity.RegisterActivityServiceServer(server, &controller.ActivityServer{})
	health.RegisterHealthServiceServer(server, &controller.HealthServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// Return the grpc server
	return server
}

// getGrpcServer creates a new grpc server
func getGrpcServer(cfg *config.Config) *grpc.Server {
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
