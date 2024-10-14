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
	var server *grpc.Server
	if cfg.Server.Mode == "OIDC" {
		// Initialise auth service & interceptor
		authSvc, err := service.NewAuthService(cfg)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to initialize auth service")
		}
		interceptor, err := service.NewAuthInterceptorService(authSvc)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to initialize interceptor")
		}

		// Create a new grpc server
		server = grpc.NewServer(
			grpc.UnaryInterceptor(interceptor.UnaryAuthMiddleware),
		)
	} else {
		server = grpc.NewServer()
	}

	// Register the service with the server
	deployment.RegisterDeploymentServiceServer(server, &controller.DeploymentServer{})
	activity.RegisterActivityServiceServer(server, &controller.ActivityServer{})
	health.RegisterHealthServiceServer(server, &controller.HealthServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// Return the grpc server
	return server
}
