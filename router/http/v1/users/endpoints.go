package users

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/domain/users"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// deploymentEndpoints is the controller for the deployment entity.
type userEndpoints struct {
	user users.Service
}

// NewDeploymentController is used to create a new deployment controller.
func newUserEndpoints(db *gorm.DB) *userEndpoints {
	return &userEndpoints{
		user: users.New(db),
	}
}

// List is used to list all users.
// @Summary List all users.
// @Security Bearer
// @Tags    🍄 User
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} ListResponse
// @Failure 500 {object} errors.Error
// @Router  /v1/users [get]
func (receiver *userEndpoints) list(ctx *gin.Context) error {
	// Get the results from the service
	results, err := receiver.user.PaginateList(ctx)
	if err != nil {
		return err
	}

	// Convert the results to the response model
	responseItems := make([]UserResponse, len(*results))
	for i, element := range *results {
		responseItems[i] = *toUserResponse(element)
	}

	// Return the response
	ctx.JSON(http.StatusOK, ListResponse{responseItems})
	return nil
}

// Read is used to read a user.
// @Summary Read a user.
// @Security Bearer
// @Tags    🍄 User
// @Accept  json
// @Produce json
// @Param   id path int true "ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router  /v1/users/{id} [get]
func (receiver *userEndpoints) read(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	// Read the user from the service
	responseItem, err := receiver.user.Read(id)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusOK, toUserResponse(responseItem))
	return nil
}

// Delete is used to delete a user.
// @Summary Delete a user.
// @Security Bearer
// @Tags    🍄 User
// @Accept  json
// @Produce json
// @Param   id path int true "ID"
// @Success 204
// @Failure 400 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router  /v1/users/{id} [delete]
func (receiver *userEndpoints) delete(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	// Delete the user using the service
	err = receiver.user.Delete(id)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// UpdateStatus is used to update a user status.
// @Summary Update a user status.
// @Security Bearer
// @Tags    🍄 User
// @Accept  json
// @Produce json
// @Param   id path int true "ID"
// @Param   status body StatusRequest true "Status"
// @Success 204
// @Failure 400 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router  /v1/users/{id}/status [put]
func (receiver *userEndpoints) UpdateStatus(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	// Parse the request
	var request StatusRequest
	if err := ctx.BindJSON(&request); err != nil {
		return err
	}

	// Update the status using the service
	statusErr := receiver.user.UpdateStatus(id, request.Status)
	if statusErr != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// Me is used to get the current user.
// @Summary Get the current user.
// @Security Bearer
// @Tags    🍄 User
// @Accept  json
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 500 {object} errors.Error
// @Router  /v1/users/me [get]
func (receiver *userEndpoints) me(ctx *gin.Context) error {
	// Get the user from the context
	user, err := receiver.user.Me(ctx)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusOK, toUserResponse(user))
	return nil
}
