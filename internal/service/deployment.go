package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/models"
	"gorm.io/gorm"
)

type Deployment struct {
	conn *gorm.DB
}

func NewDeploymentService() *Deployment {
	return &Deployment{conn: database.GetInstance()}
}

func (a *Deployment) List(ctx context.Context, page, pageSize int) ([]models.Deployment, error) {
	// Define the list of activities
	var entities []models.Deployment

	// Get the list of activities
	if err := a.conn.WithContext(ctx).
		Scopes(database.Paginate(page, pageSize)).
		Find(&entities).Error; err != nil {
		return entities, err
	}

	// Return the list of activities
	return entities, nil
}

func (a *Deployment) Get(ctx context.Context, deploymentId int) (*models.Deployment, error) {
	// Get the deployment
	var entity *models.Deployment
	if err := a.conn.WithContext(ctx).
		Where("id = ?", deploymentId).
		First(&entity).Error; err != nil {
		return nil, err
	}

	// Return the deployment
	return entity, nil
}

func (a *Deployment) Create(ctx context.Context, request models.DeploymentRequest) (*models.Deployment, error) {
	// Check if the module exists
	moduleService := NewModuleService()
	exists, err := moduleService.IsExist(ctx, request.ModuleId)
	if err != nil {
		log.Err(err).Msg("failed to check if module exists")
		return nil, fmt.Errorf("failed to check if module exists")
	}
	if !exists {
		return nil, fmt.Errorf("module does not exist")
	}

	// Convert the request to a model
	deployment := models.Deployment{
		ModuleId: request.ModuleId,
		Name:     request.Name,
		Status:   models.DeploymentStatus("pending"),
	}

	// Create the deployment to the database
	if err := a.conn.WithContext(ctx).Create(&deployment).Error; err != nil {
		return nil, err
	}

	// Add deployment variables
	if request.Variables != nil {
		for _, variable := range *request.Variables {
			variable.DeploymentID = deployment.ID
			if err := a.conn.WithContext(ctx).Create(&variable).Error; err != nil {
				// Rollback the deployment creation on failure
				a.conn.WithContext(ctx).Delete(&deployment)
				return nil, err
			}
		}
	}

	// Return the created deployment
	return &deployment, nil
}
