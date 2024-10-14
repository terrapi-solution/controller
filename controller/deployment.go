package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/deployment"
)

type DeploymentServer struct {
	rpc.DeploymentServiceServer
}

func (s *DeploymentServer) GetDeployment(ctx context.Context, req *rpc.RetrieveDeploymentRequest) (*rpc.Deployment, error) {
	deployment := &rpc.Deployment{
		Module: &rpc.DeploymentModule{},
	}

	return deployment, nil
}
