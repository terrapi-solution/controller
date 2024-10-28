package rest

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files/v2"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/terrapi-solution/controller/docs"
)

// addSwaggerRoute adds the swagger route to the router
func (r *RestServer) addSwaggerRoute(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
