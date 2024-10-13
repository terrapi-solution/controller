package server

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/activity"
)

func (s *GrpcServer) ListActivity(ctx context.Context, req *rpc.ListActivityRequest) (*rpc.Activities, error) {
	activities := &rpc.Activities{}
	return activities, nil
}

func (s *GrpcServer) InsertActivity(ctx context.Context, req *rpc.InsertActivityRequest) (*rpc.InsertActivityResponse, error) {
	response := &rpc.InsertActivityResponse{}
	return response, nil
}
