package health

import (
	"github.com/terrapi-solution/protocol/health/v1"
	"google.golang.org/grpc"
)

// RegisterService registers the activity service with the gRPC server
func RegisterService(server *grpc.Server) {
	health.RegisterHealthServiceServer(server, &GrpcHealthServer{})
}
