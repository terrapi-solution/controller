package controller

import (
	"context"
	"github.com/terrapi-solution/controller/internal/service"
	rpc "github.com/terrapi-solution/protocol/health/v1"
	"strings"
)

// HealthServer implements the HealthServiceServer interface.
type HealthServer struct {
	rpc.HealthServiceServer
	service service.HealthService
}

// Check performs a health check for the specified service.
func (s *HealthServer) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	// Create a new health service
	h := service.NewHealthService()

	// Map the service name to the corresponding health check function
	statusMap := map[string]func() rpc.CheckResponse_ServingStatus{
		"controller": h.CheckController,
		"database":   h.CheckDatabase,
	}

	checkFunc, exists := statusMap[req.Service]
	if !exists {
		// Return unknown status if the service is not recognized
		return &rpc.CheckResponse{
			Name:   strings.ToLower(req.Service),
			Status: rpc.CheckResponse_SERVING_STATUS_SERVICE_UNKNOWN,
		}, nil
	}

	return &rpc.CheckResponse{
		Name:   strings.ToLower(req.Service),
		Status: checkFunc(),
	}, nil
}

// CheckAll performs a health check for all services.
func (s *HealthServer) CheckAll(ctx context.Context, req *rpc.CheckAllRequest) (*rpc.CheckAllResponse, error) {
	// Create a new health service
	h := service.NewHealthService()

	// Create a slice of health checks
	data := rpc.CheckAllResponse{}
	data.Results = append(data.Results, &rpc.CheckResponse{Name: "controller", Status: h.CheckController()})
	data.Results = append(data.Results, &rpc.CheckResponse{Name: "database", Status: h.CheckDatabase()})

	// Return the health checks
	return &data, nil
}
