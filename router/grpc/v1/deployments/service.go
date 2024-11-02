package deployments

import (
	"github.com/terrapi-solution/protocol/deployment/v1"
	"google.golang.org/grpc"
)

// RegisterService registers the deployment service with the gRPC server
func RegisterService(server *grpc.Server) {
	deployment.RegisterDeploymentServiceServer(server, &DeploymentServer{})
}
