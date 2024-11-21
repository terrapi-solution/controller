package plans

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/router/http/errors"
)

// NewRoutesFactory creates and returns a factory to create routes
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		endpoints := newPlanEndpoints(database.GetInstance())

		modulesGroup := group.Group("/plans")
		modulesGroup.GET("", errors.HandlerWithErrorWrapper(endpoints.list))
		modulesGroup.POST("", errors.HandlerWithErrorWrapper(endpoints.add))
		modulesGroup.POST("/:id/cancel", errors.HandlerWithErrorWrapper(endpoints.cancel))
		modulesGroup.GET("/:id/variables", errors.HandlerWithErrorWrapper(endpoints.listVariable))
	}
}
