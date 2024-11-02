package deployments

import "github.com/gin-gonic/gin"

// NewRoutesFactory creates and returns a factory to create routes for deployment endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		endpoints := newDeploymentEndpoints()

		group.GET("", endpoints.list)
		group.POST("", endpoints.create)
		group.GET("/:id", endpoints.get)
		// group.PUT("/:id", endpoints.update)
		group.DELETE("/:id", endpoints.delete)
		group.GET("/:id/activities", endpoints.getActivities)
		group.GET("/:id/module", endpoints.getModule)
		group.GET("/:id/module/source", endpoints.getModuleSource)
	}
}
