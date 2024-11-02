package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/terrapi-solution/controller/docs"
	"net/http"
)

// NewRoutesFactory create and returns a factory to create routes
func NewRoutesFactory() func(group *gin.RouterGroup) {
	routesFactory := func(group *gin.RouterGroup) {
		group.GET("/swagger", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return routesFactory
}
