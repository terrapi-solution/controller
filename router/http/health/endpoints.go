package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get is used to get the health of the service.
// @Summary Get the health of the service.
// @Tags    🍏 Health
// @Accept  json
// @Produce json
// @Success 200 {object} responseHealth
// @Failure 500 {object} errors.AppError
// @Router  /health [get]
func get(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, &responseHealth{Status: "ok"})
}
