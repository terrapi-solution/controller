package plan

import (
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
)

// toPlanData converts a PlanRequest to a plan.Plan model.
func (r PlanRequest) toPlanData() plan.Plan {
	return plan.Plan{
		Name:      r.Name,
		Type:      r.Type,
		State:     plan.PendingState,
		Schedule:  r.Schedule,
		ModuleID:  r.ModuleID,
		Variables: r.toVariableData(),
	}
}

// toVariableData converts the Variables field of a PlanRequest to a slice of planVariable.PlanVariable.
func (r PlanRequest) toVariableData() []planVariable.PlanVariable {
	var variables []planVariable.PlanVariable
	for _, v := range r.Variables {
		variables = append(variables, planVariable.PlanVariable{
			Key:       v.Key,
			Value:     v.Value,
			Category:  v.Category,
			Sensitive: v.Sensitive,
		})
	}
	return variables
}
