package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/core"
	rpc "github.com/terrapi-solution/protocol/activity/v1"
)

type GrpcActivityServer struct {
	rpc.ActivityServiceServer
	Services *core.Core
}

func (s *GrpcActivityServer) List(ctx context.Context, req *rpc.ListRequest) (*rpc.ListResponse, error) {
	activities := &rpc.ListResponse{}
	return activities, nil
}

func (s *GrpcActivityServer) Insert(ctx context.Context, req *rpc.InsertRequest) (*rpc.InsertResponse, error) {

	log.Info().Msg("InsertActivity")
	log.Info().Msg(req.Message)

	response := &rpc.InsertResponse{}
	return response, nil
}
