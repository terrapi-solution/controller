package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/controller/rest"
)

func (r *RestServer) addActivityRoute(engine *gin.Engine) {
	// Create a new activity controller
	endpoints := rest.NewActivityController()

	// Create a new group for the activity route
	route := engine.Group("/v1/activities")
	route.GET("/:deploymentId", endpoints.List)
}
