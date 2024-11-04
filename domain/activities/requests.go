package activities

type ActivityRequest struct {
	DeploymentID int    `json:"deployment_id"`
	Pointer      string `json:"pointer"`
	Message      string `json:"message"`
}
