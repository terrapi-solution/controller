package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
)

// NewRoutesFactory creates and returns a factory to create routes for module endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		db := database.GetInstance()
		endpoints := newModuleEndpoints(db)

		group.GET("", endpoints.list)
		group.POST("", endpoints.create)
		group.GET("/:id", endpoints.get)
		group.DELETE("/:id", endpoints.delete)
		group.GET("/:id/source", endpoints.getSource)
		group.PUT("/:id/source", endpoints.updateSource)
	}
}
