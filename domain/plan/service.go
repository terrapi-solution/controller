package plan

import (
	"github.com/adhocore/gronx"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/plan"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
	"strings"
)

// Service struct manages interactions with activities store
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

// Add creates a new plan in the database
func (s *Service) Add(ctx *gin.Context, req PlanRequest) (plan.Plan, error) {
	// Validate the request
	err := s.AddValidation(req)
	if err != nil {
		return plan.Plan{}, err
	}

	// Convert the request to the data model
	planModel := req.toPlanData()

	// Create the plan to the database
	_, err = s.plan.Create(planModel, ctx)
	if err != nil {
		return plan.Plan{}, domainErrors.NewInternal(err, "Error creating plan", "PlanService.Add")
	}

	// Return the created plan
	return planModel, nil
}

// AddValidation validates the request to add a plan to the database
func (s *Service) AddValidation(req PlanRequest) error {
	// Validate the request
	err := s.validate.Struct(req)
	if err != nil {
		return domainErrors.NewInvalid(err, "Invalid request", "PlanService.AddValidation")
	}

	// Check if the module exists
	moduleExists := s.module.Exists(req.ModuleID)
	if !moduleExists {
		return domainErrors.NewInternal(err, "Error checking if module exists", "PlanService.AddValidation")
	}

	// Check if the plan already exists
	planExists := s.plan.ExistsByName(req.Name)
	if planExists {
		return domainErrors.NewConflict(nil, "Plan already exists", "PlanService.AddValidation")
	}

	// Check if duplicate key in variables
	if req.Variables != nil && len(req.Variables) > 0 {
		variableKeys := make(map[string]struct{})
		for _, variable := range req.Variables {
			key := strings.ToLower(variable.Key)
			if _, exists := variableKeys[key]; exists {
				return domainErrors.NewInvalid(nil, "Duplicate variable key", "PlanService.AddValidation")
			}
			variableKeys[key] = struct{}{}
		}
	}

	// Check if the schedule is valid
	if req.Type == plan.ScheduleType {
		if !gronx.New().IsValid(req.Schedule) {
			return domainErrors.NewInvalid(nil, "Invalid cron expression", "PlanService.AddValidation")
		}
	}

	return nil
}

// Cancel cancels a plan in the database
// Only plans in pending or running state can be cancelled
// TODO: Notify all workers to stop processing the plan
func (s *Service) Cancel(ctx *gin.Context, id int) error {
	// Get the plan from the database
	planModel, err := s.plan.Read(id)
	if err != nil {
		return err
	}

	// Check if the plan can be cancelled
	if planModel.State != plan.PendingState && planModel.State != plan.RunningState {
		return domainErrors.NewInvalid(nil, "Plan cannot be cancelled", "PlanService.Cancel")
	}

	// Update the plan state to cancelled in the database
	planModel.State = plan.CanceledState
	if _, err = s.plan.Update(id, planModel, ctx); err != nil {
		return domainErrors.NewInternal(err, "Error updating plan", "PlanService.Cancel")
	}

	return nil
}
