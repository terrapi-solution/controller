package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/service"
	rpc "github.com/terrapi-solution/protocol/activity/v1"
)

type GrpcActivityServer struct {
	rpc.ActivityServiceServer
}

func (s *GrpcActivityServer) List(ctx context.Context, req *rpc.ListRequest) (*rpc.ListResponse, error) {
	activities := &rpc.ListResponse{}
	return activities, nil
}

func (s *GrpcActivityServer) Insert(ctx context.Context, req *rpc.InsertRequest) (*rpc.InsertResponse, error) {

	svc := service.NewActivityService()
	create, err := svc.Create(ctx, service.ActivityRequest{
		DeploymentID: uint(req.Deployment),
		Pointer:      req.Pointer.String(),
		Message:      req.Message,
	})

	if err == nil {
		log.Debug().Msgf("Activity created with ID: %d", create.ID)
		return &rpc.InsertResponse{
			Id: int32(create.ID),
		}, nil
	}

	return nil, err
}
