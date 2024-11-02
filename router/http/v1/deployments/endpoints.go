package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/models"
	"github.com/terrapi-solution/controller/internal/services/deployment"
	"github.com/terrapi-solution/controller/internal/services/generic"
)

// deploymentEndpoints is the controller for the deployment entity.
type deploymentEndpoints struct {
	gen *generic.ServiceGeneric[models.Deployment]
	svc *deployment.ServiceDeployment
}

// NewDeploymentController is used to create a new deployment controller.
func newDeploymentEndpoints() *deploymentEndpoints {
	return &deploymentEndpoints{
		gen: generic.NewGenericService[models.Deployment](),
		svc: deployment.NewDeploymentService(),
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
// @Success 200 {object} []models.Deployment
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules [get]
func (receiver *deploymentEndpoints) list(ctx *gin.Context) { receiver.gen.List(ctx) }

// Get is used to get a specific deployment.
// @Summary Get a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   id path  int true "Deployment ID"
// @Success 200 {object} models.Deployment
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{id} [get]
func (receiver *deploymentEndpoints) get(ctx *gin.Context) { receiver.gen.GetOne(ctx) }

// Delete is used to delete a specific deployment.
// @Summary Delete a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   id path  int true "Deployment identifier"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{id} [delete]
func (receiver *deploymentEndpoints) delete(ctx *gin.Context) { receiver.gen.Delete(ctx) }

// Create is used to create a new deployment.
// @Summary Create a new deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   deployment body models.DeploymentRequest true "Deployment"
// @Success 201 {object} models.Deployment
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments [post]
func (receiver *deploymentEndpoints) create(ctx *gin.Context) {

}

// GetActivities is used to get all activities of a specific deployment.
// @Summary Get all activities of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   id           path  int true "Deployment ID"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} []models.Activity
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{id}/activities [get]
func (receiver *deploymentEndpoints) getActivities(ctx *gin.Context) {

}

// GetModule is used to get the module of a specific deployment.
// @Summary Get the module of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   id path  int true "Deployment ID"
// @Success 200 {object} models.Module
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{id}/module [get]
func (receiver *deploymentEndpoints) getModule(ctx *gin.Context) {

}

// GetModuleSource is used to get the source of the module of a specific deployment.
// @Summary Get the source of the module of a specific deployment.
// @Tags    🍑 Deployment
// @Accept  json
// @Produce json
// @Param   id path  int true "Deployment ID"
// @Success 200 {object} models.ModuleSource
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/deployments/{id}/module/source [get]
func (receiver *deploymentEndpoints) getModuleSource(ctx *gin.Context) {

}
