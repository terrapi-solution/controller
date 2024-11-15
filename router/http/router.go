package http

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/middleware/header"
	"github.com/terrapi-solution/controller/router/http/api"
	"github.com/terrapi-solution/controller/router/http/errors"
	"github.com/terrapi-solution/controller/router/http/health"
	"github.com/terrapi-solution/controller/router/http/swagger"
	"net/http"
)

// NewHttpHandler creates a new HTTP handler with all the routes defined.
func NewHttpHandler() http.Handler {
	// Set the router to release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	router := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Internal middleware
	registerInternalMiddleware(router)

	// Static routes
	registerStaticRoutes(router)

	// Base route definition
	defaultGroup := router.Group("/")
	registerRoutes(defaultGroup)

	return router
}

// registerInternalMiddleware registers custom middleware to the router.
func registerInternalMiddleware(router *gin.Engine) {
	router.Use(header.Version())
	router.Use(header.Cache())
	router.Use(header.Secure())
	router.Use(header.Options())
	router.Use(errors.Handler())
}

// registerRoutes registers all the routes to the default group.
func registerRoutes(router *gin.RouterGroup) {
	swagger.NewRoutesFactory()(router)

	// health route
	healthGroup := router.Group("/health")
	health.NewRoutesFactory()(healthGroup)

	// api route
	api.NewRoutesFactory()(router)
}

// registerStaticRoutes registers static routes to the router.
func registerStaticRoutes(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/index.css", "./static/index.css")
	router.StaticFile("/background.svg", "./static/background.svg")
}
