package rest

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/terrapi-solution/controller/controller/rest"
	_ "github.com/terrapi-solution/controller/docs"
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

	return router
}

// addHealthRoute adds the health route to the router
func (r *RestServer) addHealthRoute(engine *gin.Engine) {
	// Create a new health controller
	endpoints := rest.NewHealthController()

	// Create a new group for the health route
	deployment := engine.Group("/health")
	deployment.GET("", endpoints.Get)
}

// addSwaggerRoute adds the swagger route to the router
func (r *RestServer) addSwaggerRoute(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (r *RestServer) addActivityRoute(engine *gin.Engine) {
	// Create a new activity controller
	endpoints := rest.NewActivityController()

	// Create a new group for the activity route
	route := engine.Group("/v1/activities")
	route.GET("/:deploymentId", endpoints.List)
}

func (r *RestServer) addDeploymentRoute(engine *gin.Engine) {
	// Create a new deployment controller
	endpoints := rest.NewDeploymentController()

	// Create a new group for the deployment route
	route := engine.Group("/v1/deployments")
	route.GET("", endpoints.List)
	route.GET("/:deploymentId", endpoints.Get)
}
