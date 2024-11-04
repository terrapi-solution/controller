package deployments

import (
	"github.com/terrapi-solution/controller/data/deployment"
	"time"
)

type DeploymentResponse struct {
	ID        int                         `json:"id"`
	Name      string                      `json:"name"`
	ModuleID  int                         `json:"module_id"`
	Status    deployment.DeploymentStatus `json:"status"`
	CreatedAt time.Time                   `json:"created_at"`
}

// ListResponse struct defines books list response structure
type ListResponse struct {
	Data []DeploymentResponse `json:"data"`
}
