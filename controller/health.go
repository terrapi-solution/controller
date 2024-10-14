package controller

import (
	"context"
	"github.com/terrapi-solution/controller/internal/service"
	rpc "github.com/terrapi-solution/protocol/health"
)

// HealthServer implements the HealthServiceServer interface.
type HealthServer struct {
	rpc.HealthServiceServer
	service service.HealthService
}

// Check performs a health check for the specified service.
func (s *HealthServer) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.HealthCheck, error) {
	service := service.NewHealthService()
	statusMap := map[string]func() rpc.HealthCheck_ServingStatus{
		"controller": service.CheckController,
		"database":   service.CheckDatabase,
		"state":      service.CheckState,
	}

	checkFunc, exists := statusMap[req.Service]
	if !exists {
		// Return unknown status if the service is not recognized
		return &rpc.HealthCheck{
			Status: rpc.HealthCheck_UNKNOWN,
		}, nil
	}

	return &rpc.HealthCheck{
		Status: checkFunc(),
	}, nil
}
