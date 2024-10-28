package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/models"
	"github.com/terrapi-solution/controller/internal/services"
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

	// Get the deployment from the services
	svc := services.NewModuleService()
	deployments, err := svc.List(ctx, page, pageSize)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to list modules"))
		log.Err(err).Msg("failed to list modules")
		return
	}

	// Return the activities
	ctx.JSON(http.StatusOK, deployments)
}

// Create is used to create a new module.
// @Summary Create a new module.
// @Tags    📰 Module
// @Accept  json
// @Produce json
// @Param   module body models.ModuleRequest true "Module"
// @Success 201 {object} models.Module
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router  /v1/modules [post]
func (s *ModuleController) Create(ctx *gin.Context) {
	// Get the request
	var req models.ModuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		NewError(ctx, http.StatusBadRequest, fmt.Errorf("invalid request"))
		log.Err(err).Msg("invalid request")
		return
	}

	// Create the module
	svc := services.NewModuleService()
	module, err := svc.Create(ctx, req)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create module"))
		log.Err(err).Msg("failed to create module")
		return
	}

	// Return the module
	ctx.JSON(http.StatusCreated, module)
}
