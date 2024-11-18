package plans

import (
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
	"time"
)

type PlanResponseDto struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Type      plan.Type  `json:"type"`
	State     plan.State `json:"states"`
	Schedule  string     `json:"schedule"`
	ModuleID  int        `json:"module_id"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type PlanResponsesDto struct {
	Data []PlanResponseDto `json:"data"`
}

type PlanVariableResponseDto struct {
	ID        int                   `json:"id"`
	Key       string                `json:"key"`
	Value     string                `json:"value"`
	Category  planVariable.Category `json:"category"`
	Sensitive bool                  `json:"sensitive"`
}
