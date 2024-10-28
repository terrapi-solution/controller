package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/controller/rest"
)

func (r *RestServer) addModuleRoute(engine *gin.Engine) {
	// Create a new module controller
	endpoints := rest.NewModuleController()

	// Create a new group for the module route
	route := engine.Group("/v1/modules")
	//route.POST("", endpoints.Create)
	route.GET("", endpoints.List)
	//route.GET("/:moduleId", endpoints.Get)
	//route.DELETE("/:moduleId", endpoints.Delete)
}
