package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/service"
	"net/http"
	"strconv"
)

type DeploymentController struct{}

func NewDeploymentController() *DeploymentController {
	return &DeploymentController{}
}

// List is used to list all deployments.
// @Summary List all deployments.
// @Tags    📰 Deployment
// @Accept  json
// @Produce json
// @Param   page         query int false "Page number" default(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} database.Deployment
// @Failure 500 {object} HTTPError
// @Router  /v1/deployments [get]
func (s *DeploymentController) List(ctx *gin.Context) {
	// Get the page and page size
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	// Get the deployment from the service
	svc := service.NewDeploymentService()
	deployments, err := svc.List(ctx, page, pageSize)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to list deployments"))
		log.Err(err).Msg("failed to list deployments")
		return
	}

	// Return the activities
	ctx.JSON(http.StatusOK, deployments)
}

// Get is used to get a specific deployment.
// @Summary Get a specific deployment.
// @Tags    📰 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Success 200 {object} database.Deployment
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router  /v1/deployments/{deploymentId} [get]
func (s *DeploymentController) Get(ctx *gin.Context) {
	// Get the deployment identifier
	deploymentID, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		NewError(ctx, http.StatusBadRequest, fmt.Errorf("invalid deployment identifier"))
		log.Err(err).Msg("invalid deployment identifier")
		return
	}

	// Get the deployment from the service
	svc := service.NewDeploymentService()
	deployment, err := svc.Get(ctx, deploymentID)
	if err != nil && err.Error() == "record not found" {
		NewError(ctx, http.StatusNotFound, fmt.Errorf("deployment not found"))
		log.Err(err).Msg("deployment not found")
		return
	} else if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to get deployment"))
		log.Err(err).Msg("failed to get deployment")
		return
	}

	// Return the activities
	ctx.JSON(http.StatusOK, deployment)
}

func (s *DeploymentController) Create(ctx *gin.Context) {
	// Get the deployment from the service
	svc := service.NewDeploymentService()
	deployment, err := svc.Create(ctx)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create deployment"))
		log.Err(err).Msg("failed to create deployment")
		return
	}

	// Return the activities
	ctx.JSON(http.StatusCreated, deployment)
}
