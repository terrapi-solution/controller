package plans

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/domain/plan"
	"gorm.io/gorm"
	"net/http"
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

// List is used to list all plans.
// @Summary List all plans.
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
