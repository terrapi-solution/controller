package deployments

import "github.com/terrapi-solution/controller/data/deployment"

func toResponseModel(dpl deployment.Deployment) *DeploymentResponse {
	return &DeploymentResponse{
		ID:        dpl.ID,
		Name:      dpl.Name,
		ModuleID:  dpl.ModuleID,
		Status:    dpl.Status,
		CreatedAt: dpl.CreatedAt,
	}
}
