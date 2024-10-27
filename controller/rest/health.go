package rest

import "github.com/gin-gonic/gin"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// @BasePath /api/v1

// Get is used to delete a specific deployment.
// @Summary ping example !
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (s *HealthController) Get(c *gin.Context) {
}
