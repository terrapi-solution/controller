package grpc

import (
	v1 "github.com/terrapi-solution/controller/router/grpc/v1"
	"google.golang.org/grpc"
)

// NewGrpcService creates a new gRPC router with all the services defined.
func NewGrpcService(server *grpc.Server) {
	v1.RegisterServices()(server)
}
