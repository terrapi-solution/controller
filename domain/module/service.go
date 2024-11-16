package module

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/terrapi-solution/controller/data/module"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	db       *gorm.DB
	module   *module.Store
	validate *validator.Validate
}

// New creates a new Service struct
func New(db *gorm.DB) Service {
	return Service{
		db:       db,
		module:   module.New(db),
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Read retrieves a module entry from the database
func (s *Service) Read(id int) (module.Module, error) {
	entry, err := s.module.Read(id)
	if err != nil {
		return module.Module{}, err
	}

	return entry, nil
}

// Delete deletes a module entry to the database
func (s *Service) Delete(id int, ctx context.Context) error {
	if !s.module.Exists(id) {
		return domainErrors.NewNotFound(nil, "Module not found", "ModuleService.Delete")
	}

	if err := s.module.Delete(id, ctx); err != nil {
		return err
	}

	return nil
}

// List retrieves all module entries from the database
// This function is used only for internal purposes
func (s *Service) List() ([]module.Module, error) {
	return s.module.List()
}

// PaginateList retrieves a paginated list of module entries from the database
func (s *Service) PaginateList(ctx *gin.Context) ([]module.Module, error) {
	var entries []module.Module
	err := s.db.Model(&module.Module{}).Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search|filter.Filter),
	).Find(&entries).Error

	if err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "ModuleService.PaginateList")
	}

	return entries, nil
}

// Create creates a new module entry in the database
func (s *Service) Create(request ModuleRequest, ctx context.Context) (module.Module, error) {
	// Validate the request body
	if s.validate.Struct(request) != nil {
		return module.Module{}, domainErrors.NewInvalid(nil, "Invalid request data", "ModuleService.Create")
	}

	// Check if the module name already exists
	if s.module.ExistsByName(request.Name) {
		return module.Module{}, domainErrors.NewConflict(nil, "Module already exists with the name: "+request.Name, "ModuleService.Create")
	}

	s.db.Set("current_user", "pasunkonwn")

	// Create a new module entry in the database
	model := request.toDBModel()
	model.Type = module.UndefinedType
	entry, err := s.module.Create(model, ctx)
	if err != nil {
		return module.Module{}, err
	}

	// Return the created module
	return entry, nil
}

// SetGitConfig sets the git configuration to a module
func (s *Service) SetGitConfig(id int, request GitConfigRequest, ctx context.Context) error {
	// Validate the request data
	if err := s.validate.Struct(request); err != nil {
		return domainErrors.NewInvalid(nil, "Invalid request data", "ModuleService.SetGitConfig")
	}

	// Check if the module exists
	current, err := s.module.Read(id)
	if err != nil {
		return err
	}

	// Marshal the request to JSON
	configBytes, err := json.Marshal(request)
	if err != nil {
		return domainErrors.NewInvalid(err, "Error marshalling request data", "ModuleService.SetGitConfig")
	}

	// Update the module type and config
	current.Type = module.GitType
	current.Config = configBytes

	// Update the module in the database
	if _, err := s.module.Update(id, current, ctx); err != nil {
		return err
	}

	return nil
}

// GetGitConfig retrieves the git configuration from a module
func (s *Service) GetGitConfig(id int) (GitConfigRequest, error) {
	entry, err := s.Read(id)
	if err != nil {
		return GitConfigRequest{}, err
	}

	if entry.Type != module.GitType {
		return GitConfigRequest{}, domainErrors.NewInvalid(nil, "Module is not a git module", "ModuleService.GetGitConfig")
	}

	var config GitConfigRequest
	if err := json.Unmarshal(entry.Config, &config); err != nil {
		return GitConfigRequest{}, domainErrors.NewInvalid(err, "Error unmarshalling module config", "ModuleService.GetGitConfig")
	}

	return config, nil
}
