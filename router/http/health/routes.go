package health

import (
	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes
func NewRoutesFactory() func(group *gin.RouterGroup) {
	healthRoutesFactory := func(group *gin.RouterGroup) {
		group.GET("/", get)
	}

	return healthRoutesFactory
}
