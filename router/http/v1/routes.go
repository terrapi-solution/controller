package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/router/http/v1/activities"
	"github.com/terrapi-solution/controller/router/http/v1/deployments"
	"github.com/terrapi-solution/controller/router/http/v1/modules"
)

// NewRoutesFactory creates and returns a factory to create routes for the v1 API
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		// Create a new group for the activity route
		activityRoute := group.Group("/activities")
		activities.NewRoutesFactory()(activityRoute)

		// Create a new group for the deployment route
		deploymentRoute := group.Group("/deployments")
		deployments.NewRoutesFactory()(deploymentRoute)

		// Create a new group for the module route
		moduleRoute := group.Group("/modules")
		modules.NewRoutesFactory()(moduleRoute)
	}
}
