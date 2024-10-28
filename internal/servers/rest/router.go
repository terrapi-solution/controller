package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/terrapi-solution/controller/internal/middleware/header"
)

func (r *RestServer) loadRoute() *gin.Engine {
	// Set the router to release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	router := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Custom middleware
	router.Use(header.Version())
	router.Use(header.Cache())
	router.Use(header.Secure())
	router.Use(header.Options())

	// Route definitions
	r.addHealthRoute(router)
	r.addSwaggerRoute(router)
	r.addActivityRoute(router)
	r.addDeploymentRoute(router)
	r.addModuleRoute(router)

	return router
}
