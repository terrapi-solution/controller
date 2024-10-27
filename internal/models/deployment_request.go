package models

type DeploymentRequest struct {
	ModuleId  uint                  `json:"module_id"`
	Name      string                `json:"name"`
	Variables *[]DeploymentVariable `json:"variables"`
}
