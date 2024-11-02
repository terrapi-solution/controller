package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/models"
	"github.com/terrapi-solution/controller/internal/services/generic"
	"github.com/terrapi-solution/controller/internal/services/module"
	"net/http"
)

// moduleEndpoints is the controller for the module entity.
type moduleEndpoints struct {
	gen *generic.ServiceGeneric[models.Module]
	svc *module.Module
}

// newModuleEndpoints is used to create a new module controller.
func newModuleEndpoints() *moduleEndpoints {
	return &moduleEndpoints{
		gen: generic.NewGenericService[models.Module](),
		svc: module.NewModuleService(),
	}
}

// List is used to list all modules.
// @Summary List all modules.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} []models.Module
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules [get]
func (receiver *moduleEndpoints) list(ctx *gin.Context) { receiver.gen.List(ctx) }

// Get is used to get a specific module.
// @Summary Get a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Success 200 {object} models.Deployment
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id} [get]
func (receiver *moduleEndpoints) get(ctx *gin.Context) { receiver.gen.GetOne(ctx) }

// Delete is used to delete a specific module.
// @Summary Delete a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id} [delete]
func (receiver *moduleEndpoints) delete(ctx *gin.Context) { receiver.gen.Delete(ctx) }

// Create is used to create a new module.
// @Summary Create a new module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   module body models.ModuleRequest true "Module"
// @Success 201 {object} models.Module
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules [post]
func (receiver *moduleEndpoints) create(ctx *gin.Context) {
	// Get the request
	var req models.ModuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//NewError(ctx, http.StatusBadRequest, fmt.Errorf("invalid request"))
		log.Err(err).Msg("invalid request")
		return
	}

	// Create the module
	svc := module.NewModuleService()
	module, err := svc.Create(ctx, req)
	if err != nil {
		//NewError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create module"))
		log.Err(err).Msg("failed to create module")
		return
	}

	// Return the module
	ctx.JSON(http.StatusCreated, module)
}

// GetSource is used to get the source of a specific module.
// @Summary Get the source of a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Success 200 {object} models.ModuleSource
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id}/source [get]
func (receiver *moduleEndpoints) getSource(ctx *gin.Context) {
}

// UpdateSource is used to update the source of a specific module.
// @Summary Update the source of a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Param   source body models.ModuleSource true "Module source"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id}/source [put]
func (receiver *moduleEndpoints) updateSource(ctx *gin.Context) {
}
