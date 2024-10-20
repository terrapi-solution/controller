package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/deployment/v1"
)

type DeploymentServer struct {
	rpc.DeploymentServiceServer
}

func (s *DeploymentServer) Get(ctx context.Context, req *rpc.GetRequest) (*rpc.GetResponse, error) {
	deployment := rpc.GetResponse{
		Module: &rpc.Module{
			Name:     "hello-world",
			Address:  "https://github.com/kikitux/terraform-null-helloworld.git",
			Path:     "",
			Username: "",
			Password: "",
		},
		Request: &rpc.Request{
			Action: rpc.Action_ACTION_INIT,
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
