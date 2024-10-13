package server

import (
	"context"
	"github.com/thomas-illiet/terrapi-controller/api"
)

func (s *GrpcServer) GetDeployment(ctx context.Context, req *api.RetrieveDeploymentRequest) (*api.Deployment, error) {
	deployment := &api.Deployment{}
	deployment.Module = &api.DeploymentModule{}
	return deployment, nil
}
