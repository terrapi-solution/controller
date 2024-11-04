package activities

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
)

// NewRoutesFactory creates and returns a factory to create routes for activity endpoints
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		db := database.GetInstance()
		endpoints := newActivityEndpoints(db)

		group.GET("", endpoints.list)
		group.GET("/:id", endpoints.get)
	}
}
