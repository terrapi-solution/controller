package server

import (
	"context"
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

var authService *service.Auth

func init() {
	authService = service.NewAuthService("https://id.netboot.fr/realms/master/protocol/openid-connect/certs")
}

func (s *GrpcServer) NewGRPCServer() *grpc.Server {
	// Create a new grpc server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryServerInterceptor))

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

func UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Printf("Received request on method: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	log.Printf("Sending response from method: %s", info.FullMethod)
	_, err = authService.Validate("ddd")
	log.Printf(err.Error())
	return resp, err
}
