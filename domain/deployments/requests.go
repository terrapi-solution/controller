package deployments

type DeploymentRequest struct {
	Name      string
	ModuleID  int
	Variables []DeploymentVariableRequest
}

type DeploymentVariableRequest struct {
	Name  string
	Value string
}
