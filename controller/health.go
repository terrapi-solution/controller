package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/health"
)

type HealthServer struct {
	rpc.HealthServiceServer
}

func (s *HealthServer) Check(ctx context.Context, req *rpc.HealthCheckRequest) (*rpc.HealthCheckResponse, error) {
	deployment := &rpc.HealthCheckResponse{
		Status: rpc.HealthCheckResponse_SERVING,
	}
	return deployment, nil
}
