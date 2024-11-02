package activities

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/models"
	"github.com/terrapi-solution/controller/internal/services/generic"
)

// Activities is the controller for the activity entity.
type activityEndpoints struct {
	gen *generic.ServiceGeneric[models.Activity]
}

// NewActivityEndpoint is used creates a new activity controller
func newActivityEndpoints() *activityEndpoints {
	return &activityEndpoints{
		gen: generic.NewGenericService[models.Activity](),
	}
}

// List is used to list all activities.
// @Summary List all activity.
// @Tags    🍍 Activity
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} []models.Activity
// @Failure 500 {object} errors.AppError
// @Router  /v1/activities [get]
func (s *activityEndpoints) list(ctx *gin.Context) { s.gen.List(ctx) }

// Get is used to get a specific activity.
// @Summary Get a specific activity.
// @Tags    🍍 Activity
// @Accept  json
// @Produce json
// @Param   id path  int true "Activity identifier"
// @Success 200 {object} models.Activity
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router  /v1/activities/{id} [get]
func (s *activityEndpoints) get(ctx *gin.Context) { s.gen.GetOne(ctx) }
