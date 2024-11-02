package modules

import "github.com/gin-gonic/gin"

// NewRoutesFactory creates and returns a factory to create routes for module endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		endpoints := newModuleEndpoints()

		group.GET("", endpoints.list)
		group.POST("", endpoints.create)
		group.GET("/:id", endpoints.get)
		group.DELETE("/:id", endpoints.delete)
		group.GET("/:id/source", endpoints.getSource)
		group.PUT("/:id/source", endpoints.updateSource)
	}
}
