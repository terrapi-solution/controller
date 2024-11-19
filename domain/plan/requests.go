package plan

import (
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
)

type PlanRequest struct {
	Name      string                `json:"name" validate:"required"`
	Type      plan.Type             `json:"type" validate:"required"`
	Schedule  string                `json:"schedule"`
	ModuleID  int                   `json:"module_id" validate:"required"`
	Variables []PlanVariableRequest `json:"variables"`
}

type PlanVariableRequest struct {
	Key       string                `json:"key" validate:"required"`
	Value     string                `json:"value" validate:"required"`
	Category  planVariable.Category `json:"category" validate:"required"`
	Sensitive bool                  `json:"sensitive"`
}
