package server

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/deployment"
)

func (s *GrpcServer) GetDeployment(ctx context.Context, req *rpc.RetrieveDeploymentRequest) (*rpc.Deployment, error) {
	deployment := &rpc.Deployment{
		Module: &rpc.DeploymentModule{},
	}

	return deployment, nil
}
