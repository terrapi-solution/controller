package server

import (
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/controller/internal/service"
	"github.com/terrapi-solution/protocol/activity"
	"github.com/terrapi-solution/protocol/deployment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

type GrpcServer struct {
	// Deployment services
	Deployment *service.Deployment
	deployment.UnimplementedDeploymentServiceServer

	// Activity services
	Activity *service.Activity
	activity.UnimplementedActivityServiceServer
}

func NewGRPCServer(cfg *config.Config) *grpc.Server {
	// Initialise auth service & interceptor
	authSvc, err := service.NewAuthService(cfg)
	if err != nil {
		log.Fatalf("failed to initialize auth service: %v", err)
	}
	interceptor, err := service.NewAuthInterceptorService(authSvc)
	if err != nil {
		log.Fatalf("failed to initialize interceptor: %v", err)
	}

	// Create a new grpc server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryAuthMiddleware),
	)

	// Configure grpc instance
	srv := GrpcServer{
		Activity:   service.NewActivityService(),
		Deployment: service.NewDeploymentService(),
	}

	// Register the service with the server
	deployment.RegisterDeploymentServiceServer(server, &srv)
	activity.RegisterActivityServiceServer(server, &srv)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// Return the grpc server
	return server
}
