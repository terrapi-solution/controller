package activities

import "github.com/terrapi-solution/controller/data/activity"

func ToResponseModel(entity activity.Activity) *ActivityResponse {
	return &ActivityResponse{
		ID:           entity.ID,
		DeploymentID: entity.DeploymentID,
		Pointer:      entity.Pointer,
		Message:      entity.Message,
		CreatedAt:    entity.CreatedAt,
	}
}
