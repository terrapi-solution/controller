package activities

import "github.com/gin-gonic/gin"

// NewRoutesFactory creates and returns a factory to create routes for activity endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		endpoints := newActivityEndpoints()

		group.GET("", endpoints.list)
		group.GET("/:id", endpoints.get)
	}
}
