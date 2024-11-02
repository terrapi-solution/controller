package activities

import (
	"github.com/terrapi-solution/protocol/activity/v1"
	"google.golang.org/grpc"
)

// RegisterService registers the activity service with the gRPC server
func RegisterService(server *grpc.Server) {
	activity.RegisterActivityServiceServer(server, &GrpcActivityServer{})
}
