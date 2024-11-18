package plan

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/terrapi-solution/controller/data/plan"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	db       *gorm.DB
	module   *plan.Store
	validate *validator.Validate
}

// New creates a new Service struct
func New(db *gorm.DB) Service {
	return Service{
		db:       db,
		module:   plan.New(db),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// PaginateList retrieves a paginated list of plan entries from the database
func (s *Service) PaginateList(ctx *gin.Context) ([]plan.Plan, error) {
	var entries []plan.Plan
	err := s.db.Model(&plan.Plan{}).Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search|filter.Filter),
	).Find(&entries).Error

	if err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "PlanService.PaginateList")
	}

	return entries, nil
}
