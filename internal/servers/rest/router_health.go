package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/controller/rest"
)

// addHealthRoute adds the health route to the router
func (r *RestServer) addHealthRoute(engine *gin.Engine) {
	// Create a new health controller
	endpoints := rest.NewHealthController()

	// Create a new group for the health route
	deployment := engine.Group("/health")
	deployment.GET("", endpoints.Get)
}
