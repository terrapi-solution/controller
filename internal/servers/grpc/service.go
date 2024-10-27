package grpc

import (
	"github.com/terrapi-solution/controller/controller/grpc"
	"github.com/terrapi-solution/protocol/activity/v1"
	"github.com/terrapi-solution/protocol/deployment/v1"
	"github.com/terrapi-solution/protocol/health/v1"
)

// loadServices loads all gRPC services
func (s *GrpcServer) loadServices() {
	deployment.RegisterDeploymentServiceServer(s.server, &grpc.DeploymentServer{})
	activity.RegisterActivityServiceServer(s.server, &grpc.GrpcActivityServer{})
	health.RegisterHealthServiceServer(s.server, &grpc.GrpcHealthServer{})
}
