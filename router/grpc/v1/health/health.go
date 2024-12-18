package health

import (
	"context"
	"github.com/terrapi-solution/controller/internal/core"
	"github.com/terrapi-solution/controller/internal/services/health"
	rpc "github.com/terrapi-solution/protocol/health/v1"
	"strings"
)

// GrpcHealthServer implements the HealthServiceServer interface.
type GrpcHealthServer struct {
	rpc.HealthServiceServer
	Services *core.Core
	service  health.HealthService
}

// Check performs a health check for the specified services.
func (s *GrpcHealthServer) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	// Create a new health services
	//h := services.NewHealthService()

	// Map the services name to the corresponding health check function
	//statusMap := map[string]func() rpc.CheckResponse_ServingStatus{
	//	"controller": h.CheckController,
	//	"database":   h.CheckDatabase,
	//}

	//checkFunc, exists := statusMap[req.Service]
	//if !exists {
	//	// Return unknown status if the services is not recognized
	//	return &rpc.CheckResponse{
	//		Name:   strings.ToLower(req.Service),
	//		Status: rpc.CheckResponse_SERVING_STATUS_SERVICE_UNKNOWN,
	//	}, nil
	//}

	return &rpc.CheckResponse{
		Name: strings.ToLower(req.Service),
		//Status: checkFunc(),
	}, nil
}

// CheckAll performs a health check for all services.
func (s *GrpcHealthServer) CheckAll(ctx context.Context, req *rpc.CheckAllRequest) (*rpc.CheckAllResponse, error) {
	// Create a new health services
	//h := services.NewHealthService()

	// Create a slice of health checks
	data := rpc.CheckAllResponse{}
	//data.Results = append(data.Results, &rpc.CheckResponse{Name: "controller", Status: h.CheckController()})
	//data.Results = append(data.Results, &rpc.CheckResponse{Name: "database", Status: h.CheckDatabase()})

	// Return the health checks
	return &data, nil
}
