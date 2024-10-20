package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/activity/v1"
)

type ActivityServer struct {
	rpc.ActivityServiceServer
}

func (s *ActivityServer) ListActivity(ctx context.Context, req *rpc.ListRequest) (*rpc.ListResponse, error) {
	activities := &rpc.ListResponse{}
	return activities, nil
}

func (s *ActivityServer) InsertActivity(ctx context.Context, req *rpc.InsertRequest) (*rpc.InsertResponse, error) {
	response := &rpc.InsertResponse{}
	return response, nil
}
