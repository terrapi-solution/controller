package v1

import (
	"github.com/terrapi-solution/controller/router/grpc/v1/activities"
	"github.com/terrapi-solution/controller/router/grpc/v1/deployments"
	"github.com/terrapi-solution/controller/router/grpc/v1/health"
	"google.golang.org/grpc"
)

func RegisterServices() func(server *grpc.Server) {
	return func(server *grpc.Server) {
		health.RegisterService(server)
		deployments.RegisterService(server)
		activities.RegisterService(server)
	}
}
