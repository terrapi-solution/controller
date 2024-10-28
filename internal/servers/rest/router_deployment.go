package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/controller/rest"
)

func (r *RestServer) addDeploymentRoute(engine *gin.Engine) {
	// Create a new deployment controller
	endpoints := rest.NewDeploymentController()

	// Create a new group for the deployment route
	route := engine.Group("/v1/deployments")
	route.GET("", endpoints.List)
	route.POST("", endpoints.Create)
	route.GET("/:deploymentId", endpoints.Get)

}
