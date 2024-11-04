package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/domain/modules"
	"gorm.io/gorm"
)

// moduleEndpoints is the controller for the module entity.
type moduleEndpoints struct {
	svc *modules.Service
}

// newModuleEndpoints is used to create a new module controller.
func newModuleEndpoints(db *gorm.DB) *moduleEndpoints {
	return &moduleEndpoints{
		svc: modules.New(db),
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
// @Success 200 {object} []ModuleResponse
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules [get]
func (receiver *moduleEndpoints) list(ctx *gin.Context) {
	//receiver.gen.List(ctx)
}

// Get is used to get a specific module.
// @Summary Get a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Success 200 {object} ModuleResponse
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id} [get]
func (receiver *moduleEndpoints) get(ctx *gin.Context) {
	//receiver.gen.GetOne(ctx)
}

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
func (receiver *moduleEndpoints) delete(ctx *gin.Context) {
	//receiver.gen.Delete(ctx)
}

// Create is used to create a new module.
// @Summary Create a new module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   module body ModuleResponse true "Module"
// @Success 201 {object} ModuleResponse
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules [post]
func (receiver *moduleEndpoints) create(ctx *gin.Context) {

}

// GetSource is used to get the source of a specific module.
// @Summary Get the source of a specific module.
// @Tags    🍆 Module
// @Accept  json
// @Produce json
// @Param   id path  int true "Module identifier"
// @Success 200 {object} ModuleResponse
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
// @Param   source body ModuleResponse true "Module source"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/modules/{id}/source [put]
func (receiver *moduleEndpoints) updateSource(ctx *gin.Context) {
}
