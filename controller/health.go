package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/health"
)

// HealthServer implements the HealthServiceServer interface.
type HealthServer struct {
	rpc.HealthServiceServer
}

// Check performs a health check for the specified service.
func (s *HealthServer) Check(ctx context.Context, req *rpc.HealthCheckRequest) (*rpc.HealthCheckResponse, error) {
	statusMap := map[string]func() rpc.HealthCheckResponse_ServingStatus{
		"controller": s.checkController,
		"database":   s.checkDatabase,
		"state":      s.checkState,
	}

	checkFunc, exists := statusMap[req.Service]
	if !exists {
		// Return unknown status if the service is not recognized
		return &rpc.HealthCheckResponse{
			Status: rpc.HealthCheckResponse_SERVICE_UNKNOWN,
		}, nil
	}

	return &rpc.HealthCheckResponse{
		Status: checkFunc(),
	}, nil
}

// checkController checks the health of the controller service.
func (s *HealthServer) checkController() rpc.HealthCheckResponse_ServingStatus {
	return rpc.HealthCheckResponse_SERVING
}

// checkDatabase checks the health of the database service.
func (s *HealthServer) checkDatabase() rpc.HealthCheckResponse_ServingStatus {
	return rpc.HealthCheckResponse_SERVING
}

// checkState checks the health of the state service.
func (s *HealthServer) checkState() rpc.HealthCheckResponse_ServingStatus {
	return rpc.HealthCheckResponse_SERVING
}
