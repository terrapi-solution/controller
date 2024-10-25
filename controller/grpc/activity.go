package grpc

import (
	"context"
	"github.com/terrapi-solution/controller/internal/core"
	rpc "github.com/terrapi-solution/protocol/activity/v1"
)

type GrpcActivityServer struct {
	rpc.ActivityServiceServer
	Services *core.Core
}

func (s *GrpcActivityServer) ListActivity(ctx context.Context, req *rpc.ListRequest) (*rpc.ListResponse, error) {
	activities := &rpc.ListResponse{}
	return activities, nil
}

func (s *GrpcActivityServer) InsertActivity(ctx context.Context, req *rpc.InsertRequest) (*rpc.InsertResponse, error) {
	response := &rpc.InsertResponse{}
	return response, nil
}
