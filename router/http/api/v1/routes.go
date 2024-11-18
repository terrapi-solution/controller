package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/router/http/api/v1/modules"
	"github.com/terrapi-solution/controller/router/http/api/v1/plans"
)

// NewRoutesFactory creates and returns a factory to create routes for the v1 API
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		v1Group := group.Group("/v1")
		modules.NewRoutesFactory()(v1Group)
		plans.NewRoutesFactory()(v1Group)
	}
}
