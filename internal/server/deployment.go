package server

import (
	"context"
	"github.com/thomas-illiet/terrapi-controller/api"
)

func (s *GrpcServer) GetDeployment(ctx context.Context, req *api.RetrieveDeploymentRequest) (*api.Deployment, error) {
	deployment := &api.Deployment{}
	deployment.Module = &api.DeploymentModule{}
	deployment.Module.Id = req.Id
	return deployment, nil
}

func (s *GrpcServer) ListDeployment(ctx context.Context, req *api.ListDeploymentRequest) (*api.Deployments, error) {
	deployments := &api.Deployments{}
	return deployments, nil
}
