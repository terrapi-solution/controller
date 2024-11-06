package users

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/router/http/errors"
)

// NewRoutesFactory creates and returns a factory to create routes
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {
		db := database.GetInstance()
		endpoints := newUserEndpoints(db)

		group.GET("", errors.HandlerWithErrorWrapper(endpoints.list))
		group.GET("/me", errors.HandlerWithErrorWrapper(endpoints.me))
		group.GET("/:id", errors.HandlerWithErrorWrapper(endpoints.read))
		group.DELETE("/:id", errors.HandlerWithErrorWrapper(endpoints.delete))
		group.PUT("/:id/status", errors.HandlerWithErrorWrapper(endpoints.UpdateStatus))
	}
}
