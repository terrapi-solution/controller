package server

import (
	"context"
	"github.com/thomas-illiet/terrapi-controller/api"
)

func (s *GrpcServer) ListActivity(ctx context.Context, req *api.ListActivityRequest) (*api.Activities, error) {
	activities := &api.Activities{}
	return activities, nil
}

func (s *GrpcServer) InsertActivity(ctx context.Context, req *api.InsertActivityRequest) (*api.InsertActivityResponse, error) {
	response := &api.InsertActivityResponse{}
	return response, nil
}
