package controller

import (
	"context"
	rpc "github.com/terrapi-solution/protocol/activity"
)

type ActivityServer struct {
	rpc.ActivityServiceServer
}

func (s *ActivityServer) ListActivity(ctx context.Context, req *rpc.ListActivityRequest) (*rpc.Activities, error) {
	activities := &rpc.Activities{}
	return activities, nil
}

func (s *ActivityServer) InsertActivity(ctx context.Context, req *rpc.InsertActivityRequest) (*rpc.InsertActivityResponse, error) {
	response := &rpc.InsertActivityResponse{}
	return response, nil
}
