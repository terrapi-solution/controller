package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/deployment"
)

type DeploymentServer struct {
	rpc.DeploymentServiceServer
}

func (s *DeploymentServer) Get(ctx context.Context, req *rpc.RetrieveRequest) (*rpc.Deployment, error) {
	deployment := rpc.Deployment{
		Module: &rpc.Module{
			Name:     "hello-world",
			Address:  "https://github.com/kikitux/terraform-null-helloworld.git",
			Path:     "",
			Username: "",
			Password: "",
		},
		Request: &rpc.Request{
			Action: rpc.Request_init,
			Variables: []*rpc.RequestVariable{
				{
					Name:   "name",
					Value:  "dd",
					Secret: true,
				},
			},
		},
	}

	return &deployment, nil
}
