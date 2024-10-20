package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/activity/v1"
)

type GrpcActivityServer struct {
	rpc.ActivityServiceServer
}

func (s *GrpcActivityServer) ListActivity(ctx context.Context, req *rpc.ListRequest) (*rpc.ListResponse, error) {
	activities := &rpc.ListResponse{}
	return activities, nil
}

func (s *GrpcActivityServer) InsertActivity(ctx context.Context, req *rpc.InsertRequest) (*rpc.InsertResponse, error) {
	response := &rpc.InsertResponse{}
	return response, nil
}
