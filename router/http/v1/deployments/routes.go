package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
)

// NewRoutesFactory creates and returns a factory to create routes for deployment endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		db := database.GetInstance()
		endpoints := newDeploymentEndpoints(db)

		group.GET("", endpoints.list)
		group.POST("", endpoints.create)
		group.GET("/:deploymentId", endpoints.get)
		// group.PUT("/:id", endpoints.update)
		group.DELETE("/:deployments", endpoints.delete)
		group.GET("/:deploymentId/activities", endpoints.getActivities)
		group.GET("/:deploymentId/module", endpoints.getModule)
		group.GET("/:deploymentId/module/source", endpoints.getModuleSource)
	}
}
