package plans

import (
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
)

func toPlanDto(entity plan.Plan) PlanResponseDto {
	return PlanResponseDto{
		ID:        entity.ID,
		Name:      entity.Name,
		Type:      entity.Type,
		State:     entity.State,
		Schedule:  entity.Schedule,
		ModuleID:  entity.ModuleID,
		CreatedBy: entity.CreatedBy.CreatedBy,
		UpdatedBy: entity.UpdatedBy.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func PlanVariableDto(entity planVariable.PlanVariable) PlanVariableResponseDto {
	// Remove sensitive data from the response
	var dataValue string
	if entity.Sensitive {
		dataValue = "********"
	} else {
		dataValue = entity.Value
	}

	// Return the response
	return PlanVariableResponseDto{
		ID:        entity.ID,
		Key:       entity.Key,
		Value:     dataValue,
		Category:  entity.Category,
		Sensitive: entity.Sensitive,
	}
}
