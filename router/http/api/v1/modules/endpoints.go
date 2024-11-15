package modules

import (
	"github.com/gin-gonic/gin"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/domain/module"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// deploymentEndpoints is the controller for the deployment entity.
type moduleEndpoints struct {
	svc module.Service
}

// NewDeploymentController is used to create a new deployment controller.
func newModuleEndpoints(db *gorm.DB) *moduleEndpoints {
	return &moduleEndpoints{
		svc: module.New(db),
	}
}

// List is used to list all modules.
// @Summary List all modules.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} ListResponseDto
// @Router  /api/v1/modules [get]
func (receiver *moduleEndpoints) list(ctx *gin.Context) error {
	// Get the results from the service
	results, err := receiver.svc.PaginateList(ctx)
	if err != nil {
		return err
	}

	// Convert the results to the response model
	responseItems := make([]ModuleResponseDto, len(results))
	for i, element := range results {
		responseItems[i] = toModuleDto(element)
	}

	// Return the response
	ctx.JSON(http.StatusOK, ListResponseDto{responseItems})
	return nil
}

// read is used to get a module by id.
// @Summary Get a module by id.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   id path string true "Module ID"
// @Success 200 {object} ModuleResponseDto
// @Failure 404 {object} errors.Error
// @Router  /api/v1/modules/{id} [get]
func (receiver *moduleEndpoints) read(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleRoute.Read")
	}

	// Get the result from the service
	result, err := receiver.svc.Read(id)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusOK, toModuleDto(result))
	return nil
}

// create is used to create a new module.
// @Summary Create a new module.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   body body module.ModuleRequest true "Module"
// @Success 201 {object} ModuleResponseDto
// @Failure 400 {object} errors.Error
// @Router  /api/v1/modules [post]
func (receiver *moduleEndpoints) create(ctx *gin.Context) error {
	// Parse the request body
	var request module.ModuleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return domainErrors.NewInvalid(err, "Invalid request body", "ModuleRoute.Create")
	}

	// Create the module
	result, err := receiver.svc.Create(request, ctx)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusCreated, toModuleDto(result))
	return nil
}

// delete is used to delete a module by id.
// @Summary Delete a module by id.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   id path string true "Module ID"
// @Success 204
// @Failure 404 {object} errors.Error
// @Router  /api/v1/modules/{id} [delete]
func (receiver *moduleEndpoints) delete(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleRoute.Delete")
	}

	err = receiver.svc.Delete(id, ctx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// setGitConfig is used to define a git configuration to a module.
// @Summary Set git configuration to a module.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   id path string true "Module ID"
// @Param   body body module.GitConfigRequest true "Git Configuration"
// @Success 204 {object} GitConfigDto
// @Failure 404 {object} errors.Error
// @Router  /api/v1/modules/{id}/config/git [post]
func (receiver *moduleEndpoints) setGitConfig(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleRoute.SetGitConfig")
	}

	// Parse the request body
	var request module.GitConfigRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return domainErrors.NewInvalid(err, "Invalid request body", "ModuleRoute.SetGitConfig")
	}

	// Create the module
	if err := receiver.svc.SetGitConfig(id, request, ctx); err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusOK, toGitConfigDto(request))
	return nil
}

// getGitConfig is used to get the git configuration of a module.
// @Summary Get git configuration of a module.
// @Security Bearer
// @Tags    🍄 Module
// @Accept  json
// @Produce json
// @Param   id path string true "Module ID"
// @Success 200 {object} GitConfigDto
// @Failure 404 {object} errors.Error
// @Router  /api/v1/modules/{id}/config/git [get]
func (receiver *moduleEndpoints) getGitConfig(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleRoute.Read")
	}

	// Get the result from the service
	result, err := receiver.svc.GetGitConfig(id)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusOK, toGitConfigDto(result))
	return nil
}
