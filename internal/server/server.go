package server

import (
	"github.com/thomas-illiet/terrapi-controller/api"
	"github.com/thomas-illiet/terrapi-controller/internal/service"
	"google.golang.org/grpc"
	"log"
)

type GrpcServer struct {
	api.UnimplementedDeploymentServiceServer
	api.UnimplementedActivityServiceServer
	Activity   *service.Activity
	Deployment *service.Deployment
}

func (s *GrpcServer) NewGRPCServer() *grpc.Server {
	// Initialise our auth service & interceptor
	authSvc, err := service.NewAuthService("https://id.netboot.fr/realms/master")
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
	api.RegisterDeploymentServiceServer(server, &srv)
	api.RegisterActivityServiceServer(server, &srv)

	// Return the grpc server
	return server
}
