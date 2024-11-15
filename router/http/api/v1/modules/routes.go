package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/router/http/errors"
)

// NewRoutesFactory creates and returns a factory to create routes
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		endpoints := newModuleEndpoints(database.GetInstance())

		modulesGroup := group.Group("/modules")
		modulesGroup.GET("", errors.HandlerWithErrorWrapper(endpoints.list))
		modulesGroup.POST("", errors.HandlerWithErrorWrapper(endpoints.create))
		modulesGroup.DELETE("", errors.HandlerWithErrorWrapper(endpoints.delete))
		modulesGroup.GET(":id", errors.HandlerWithErrorWrapper(endpoints.read))
		modulesGroup.GET(":id/config/git", errors.HandlerWithErrorWrapper(endpoints.getGitConfig))
		modulesGroup.POST(":id/config/git", errors.HandlerWithErrorWrapper(endpoints.setGitConfig))
	}
}
