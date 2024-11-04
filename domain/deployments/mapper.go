package deployments

import (
	"github.com/terrapi-solution/controller/data/deployment"
	"github.com/terrapi-solution/controller/data/deploymentVariable"
)

// toDeploymentDBModel converts DeploymentRequest to Deployment
func toDeploymentDBModel(request DeploymentRequest) *deployment.Deployment {
	return &deployment.Deployment{
		Name:     request.Name,
		ModuleID: request.ModuleID,
		Status:   deployment.Unknown,
	}
}

// toDeploymentVariableDBModel converts DeploymentVariableRequest to DeploymentVariable
func toDeploymentVariableDBModel(request DeploymentVariableRequest, deploymentId int) *deploymentVariable.DeploymentVariable {
	return &deploymentVariable.DeploymentVariable{
		Name:         request.Name,
		Value:        request.Value,
		DeploymentID: deploymentId,
	}
}
