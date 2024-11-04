package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/domain/deployments"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/router/http/v1/activities"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// deploymentEndpoints is the controller for the deployment entity.
type deploymentEndpoints struct {
	svc deployments.Service
}

// NewDeploymentController is used to create a new deployment controller.
func newDeploymentEndpoints(db *gorm.DB) *deploymentEndpoints {
	return &deploymentEndpoints{
		svc: deployments.New(db),
	}
}

// List is used to list all deployments.
// @Summary List all deployments.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} ListResponse
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments [get]
func (receiver *deploymentEndpoints) list(ctx *gin.Context) {
	// Get the results from the service
	results, err := receiver.svc.List(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Convert the results to the response model
	responseItems := make([]DeploymentResponse, len(results))
	for i, element := range results {
		responseItems[i] = *toResponseModel(element)
	}

	// Return the response
	ctx.JSON(http.StatusOK, ListResponse{Data: responseItems})
}

// Get is used to get a specific deployment.
// @Summary Get a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Success 200 {object} DeploymentResponse
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{deploymentId} [get]
func (receiver *deploymentEndpoints) get(ctx *gin.Context) {
	deploymentId, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.NotFound))
		return
	}

	// Get the results from the service
	result, err := receiver.svc.ReadFirst(deploymentId)
	if err != nil && err.(*domainErrors.AppError) != nil {
		_ = ctx.Error(err)
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, toResponseModel(*result))
}

// Delete is used to delete a specific deployment.
// @Summary Delete a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment identifier"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{deploymentId} [delete]
func (receiver *deploymentEndpoints) delete(ctx *gin.Context) {
	deploymentId, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.NotFound))
		return
	}

	// Get the results from the service
	dellErr := receiver.svc.Delete(deploymentId)
	if dellErr != nil && dellErr.(*domainErrors.AppError) != nil {
		_ = ctx.Error(err)
		return
	}

	// Return the response
	ctx.JSON(http.StatusNoContent, nil)

}

// Create is used to create a new deployment.
// @Summary Create a new deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deployment body deployments.DeploymentRequest true "Deployment"
// @Success 201 {object} DeploymentResponse
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments [post]
func (receiver *deploymentEndpoints) create(ctx *gin.Context) {
	var request deployments.DeploymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.ValidationError))
		return
	}

	// Get the results from the service
	result, err := receiver.svc.Create(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	// Return the response
	ctx.JSON(http.StatusCreated, toResponseModel(*result))
}

// GetActivities is used to get all activities of a specific deployment.
// @Summary Get all activities of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} []activities.ActivityResponse
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{deploymentId}/activities [get]
func (receiver *deploymentEndpoints) getActivities(ctx *gin.Context) {
	deploymentId, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.NotFound))
		return
	}

	// Get the results from the service
	// TODO: Improve error handling mechanism
	results, _ := receiver.svc.ListActivities(deploymentId, ctx)
	//if err != nil {
	//	_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.UnknownError))
	//	return
	//}

	// Convert the results to the response model
	responseItems := make([]activities.ActivityResponse, len(results))
	for i, element := range results {
		responseItems[i] = *activities.ToResponseModel(element)
	}

	// Return the response
	ctx.JSON(http.StatusOK, activities.ListResponse{Data: responseItems})
}

// GetModule is used to get the module of a specific deployment.
// @Summary Get the module of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Success 200 {object} errors.AppError
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{deploymentId}/module [get]
func (receiver *deploymentEndpoints) getModule(ctx *gin.Context) {
	_, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.NotFound))
		return
	}
}

// GetModuleSource is used to get the source of the module of a specific deployment.
// @Summary Get the source of the module of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deploymentId path  int true "Deployment ID"
// @Success 200 {object} errors.AppError
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{deploymentId}/module/source [get]
func (receiver *deploymentEndpoints) getModuleSource(ctx *gin.Context) {
	_, err := strconv.Atoi(ctx.Param("deploymentId"))
	if err != nil {
		_ = ctx.Error(domainErrors.NewAppErrorWithType(domainErrors.NotFound))
		return
	}
}
