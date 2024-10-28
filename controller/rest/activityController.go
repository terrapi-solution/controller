package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/services"
	"net/http"

	"strconv"
)

type ActivityController struct {
}

// NewActivityController creates a new activity controller
func NewActivityController() *ActivityController {
	return &ActivityController{}
}

// List is used to list all deployments for a specific deployment.
// @Summary List all deployments for a specific deployment
// @Tags    📰 Activity
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Param   page         query int false "Page number" default(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} models.Activity
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router  /v1/activities/{deploymentId} [get]
func (s *ActivityController) List(ctx *gin.Context) {
	// Get the deployment identifier
	deploymentID, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		NewError(ctx, http.StatusBadRequest, fmt.Errorf("invalid deployment identifier"))
		log.Err(err).Msg("invalid deployment identifier")
		return
	}

	// Get the page and page size
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))

	// Get the activities from the services
	svc := services.NewActivityService()
	activities, err := svc.List(ctx, deploymentID, page, pageSize)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to list activities"))
		log.Err(err).Msg("failed to list activities")
		return
	}
	if len(activities) == 0 {
		NewError(ctx, http.StatusNotFound, fmt.Errorf("no activities found for requested deployment"))
		return
	}

	// Return the activities
	ctx.JSON(http.StatusOK, activities)
}
