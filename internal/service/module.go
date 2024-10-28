package service

import (
	"context"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/models"
	"gorm.io/gorm"
)

type Module struct {
	conn *gorm.DB
}

func NewModuleService() *Module {
	return &Module{conn: database.GetInstance()}
}

// IsExist is used to check if a module exists.
func (s *Module) IsExist(ctx context.Context, moduleId uint) (bool, error) {
	var count int64
	err := s.conn.WithContext(ctx).Model(&Module{}).Where("id = ?", moduleId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// List is used to get a list of modules.
func (s *Module) List(ctx context.Context, page, pageSize int) ([]Module, error) {
	var entities []Module
	if err := s.conn.WithContext(ctx).
		Scopes(database.Paginate(page, pageSize)).
		Find(&entities).Error; err != nil {
		return entities, err
	}
	return entities, nil
}

// Get is used to get a module by ID.
func (s *Module) Get(ctx context.Context, moduleId uint) (*models.Module, error) {
	var entity *models.Module
	if err := s.conn.WithContext(ctx).
		Where("id = ?", moduleId).
		First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// Create is used to create a module.
func (s *Module) Create(ctx context.Context, request models.ModuleRequest) (*models.Module, error) {
	// Convert the request to a model
	entity := &models.Module{
		Name: request.Name,
	}

	// Create the module to the database
	if err := s.conn.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}

	// Add the source to the module
	if err := s.conn.WithContext(ctx).Create(&entity.Source).Error; err != nil {
		// Rollback the deployment creation on failure
		s.conn.WithContext(ctx).Delete(&entity)
		return nil, err
	}

	return entity, nil
}
