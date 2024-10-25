package rest

import "github.com/gin-gonic/gin"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// DeploymentDelete is used to delete a specific deployment.
func (s *HealthController) Get(c *gin.Context) {
}
