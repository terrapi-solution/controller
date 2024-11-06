package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/router/http/v1/users"
)

// NewRoutesFactory creates and returns a factory to create routes for the v1 API
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		// users route
		v1UsersGroup := group.Group("/users")
		users.NewRoutesFactory()(v1UsersGroup)

	}
}
