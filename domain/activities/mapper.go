package activities

import "github.com/terrapi-solution/controller/data/activity"

func toActivityDBModel(request ActivityRequest) *activity.Activity {
	return &activity.Activity{
		DeploymentID: request.DeploymentID,
		Pointer:      request.Pointer,
		Message:      request.Message,
	}
}
