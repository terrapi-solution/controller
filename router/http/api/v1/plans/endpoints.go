package plans

import (
	"github.com/gin-gonic/gin"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/domain/plan"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// deploymentEndpoints is the controller for the deployment entity.
type planEndpoints struct {
	svc plan.Service
}

// newPlanEndpoints is used to create a new plan controller.
func newPlanEndpoints(db *gorm.DB) *planEndpoints {
	return &planEndpoints{
		svc: plan.New(db),
	}
}

// List is used to list all execution plans.
// @Summary List all execution plans.
// @Security Bearer
// @Tags 🍑 Plans
// @Accept  json
// @Produce json
// @Param   search       query string false "Search"
// @Param   filter       query []string false "Filter"
// @Param   page         query int false "Page" default(1) minimum(1)
// @Param   page_size    query int false "Page size" default(10) minimum(1) maximum(100)
// @Param   order_by     query string false "Order by" default(id)
// @Param   order_direction query string false "Order direction" default(desc) enum(desc,asc)
// @Success 200 {object} PlanResponsesDto
// @Router  /api/v1/plans [get]
func (receiver *planEndpoints) list(ctx *gin.Context) error {
	// Get the results from the service
	results, err := receiver.svc.PaginateList(ctx)
	if err != nil {
		return err
	}

	// Convert the results to the response model
	responseItems := make([]PlanResponseDto, len(results))
	for i, element := range results {
		responseItems[i] = toPlanDto(element)
	}

	// Return the response
	ctx.JSON(http.StatusOK, PlanResponsesDto{Data: responseItems})
	return nil
}

// Add is used to create a new execution plan.
// @Summary Create a new execution plan.
// @Security Bearer
// @Tags 🍑 Plans
// @Accept  json
// @Produce json
// @Param   request body plan.PlanRequest true "Request"
// @Success 201 {object} PlanResponseDto
// @Router  /api/v1/plans [post]
func (receiver *planEndpoints) add(ctx *gin.Context) error {
	// Parse the request
	var request plan.PlanRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return err
	}

	// Create the plan
	result, err := receiver.svc.Add(ctx, request)
	if err != nil {
		return err
	}

	// Return the response
	ctx.JSON(http.StatusCreated, toPlanDto(result))
	return nil
}

// cancel is used to cancel an execution plan.
// @Summary Cancel an execution plan.
// @Security Bearer
// @Tags 🍑 Plans
// @Accept  json
// @Produce json
// @Param   id path string true "Plan ID"
// @Success 204
// @Failure 404 {object} errors.Error
// @Router  /api/v1/plans/{id}/cancel [post]
func (receiver *planEndpoints) cancel(ctx *gin.Context) error {
	// Get the ID from the URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleRoute.Read")
	}

	// Cancel the plan
	err = receiver.svc.Cancel(ctx, id)
	if err != nil {
		return err
	}

	// Return the response
	ctx.Status(http.StatusNoContent)
	return nil
}
