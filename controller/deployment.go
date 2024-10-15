package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/deployment"
)

type DeploymentServer struct {
	rpc.DeploymentServiceServer
}

func (s *DeploymentServer) GetDeployment(ctx context.Context, req *rpc.RetrieveRequest) (*rpc.Deployment, error) {
	deployment := &rpc.Deployment{
		Module: &rpc.Module{},
	}

	return deployment, nil
}
