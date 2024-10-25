package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/controller/rest"
	"github.com/terrapi-solution/controller/internal/middleware/header"
)

func (s *RestServer) loadRoute() *gin.Engine {
	// Set the router to release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	r := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Custom middleware
	r.Use(header.Version())
	r.Use(header.Cache())
	r.Use(header.Secure())
	r.Use(header.Options())

	// Route definitions
	s.addHealthRoute(r)

	return r
}

// addHealthRoute adds the health route to the router
func (s *RestServer) addHealthRoute(engine *gin.Engine) {
	// Create a new health controller
	endpoints := rest.NewHealthController()

	// Create a new group for the health route
	deployment := engine.Group("/health")
	deployment.GET("", endpoints.Get)
}
