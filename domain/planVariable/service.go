package planVariable

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
)

// Service struct manages interactions with execution plan variable store
type Service struct {
	db       *gorm.DB
	plan     *plan.Store
	module   *module.Store
	validate *validator.Validate
}

// New creates a new Service struct
func New(db *gorm.DB) Service {
	return Service{
		db:       db,
		plan:     plan.New(db),
		module:   module.New(db),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// PaginateList retrieves a paginated list of plan variable entries from the database
func (s *Service) PaginateList(planId int, ctx *gin.Context) ([]planVariable.PlanVariable, error) {
	var entries []planVariable.PlanVariable
	err := s.db.Model(&planVariable.PlanVariable{}).Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search|filter.Filter),
	).Where("plan_id = ?", planId).Find(&entries).Error

	if err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "PlanVariableService.PaginateList")
	}

	return entries, nil
}
