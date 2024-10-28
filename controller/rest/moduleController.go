package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/service"
	"net/http"
	"strconv"
)

type ModuleController struct{}

func NewModuleController() *ModuleController {
	return &ModuleController{}
}

// List is used to list all modules.
// @Summary List all modules.
// @Tags    📰 Module
// @Accept  json
// @Produce json
// @Param   page         query int false "Page number" default(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} []models.Module
// @Failure 500 {object} HTTPError
// @Router  /v1/modules [get]
func (s *ModuleController) List(ctx *gin.Context) {
	// Get the page and page size
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	// Get the deployment from the service
	svc := service.NewModuleService()
	deployments, err := svc.List(ctx, page, pageSize)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to list modules"))
		log.Err(err).Msg("failed to list modules")
		return
	}

	// Return the activities
	ctx.JSON(http.StatusOK, deployments)
}
