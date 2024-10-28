package service

import (
	"context"
	"errors"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/models"
)

type Deployment struct {
}

func NewDeploymentService() *Deployment {
	return &Deployment{}
}

func (a *Deployment) List(ctx context.Context, page, pageSize int) ([]models.Deployment, error) {
	// Define the list of activities
	var entities []models.Deployment

	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return entities, errors.New("database instance is not initialized")
	}

	// Get the list of activities

	if err := conn.WithContext(ctx).
		Scopes(database.Paginate(page, pageSize)).
		Find(&entities).Error; err != nil {
		return entities, err
	}

	// Return the list of activities
	return entities, nil
}

func (a *Deployment) Get(ctx context.Context, deploymentId int) (*models.Deployment, error) {
	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return nil, errors.New("database instance is not initialized")
	}

	// Get the deployment
	var entity *models.Deployment
	if err := conn.WithContext(ctx).
		Where("id = ?", deploymentId).
		First(&entity).Error; err != nil {
		return nil, err
	}

	// Return the deployment
	return entity, nil
}

func (a *Deployment) Create(ctx context.Context, request models.DeploymentRequest) (*models.Deployment, error) {
	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return nil, errors.New("database instance is not initialized")
	}

	// Convert the request to a model
	deployment := models.Deployment{
		ModuleId: request.ModuleId,
		Name:     request.Name,
		Status:   models.DeploymentStatus("pending"),
	}

	// Create the deployment to the database
	if err := conn.WithContext(ctx).Create(&deployment).Error; err != nil {
		return nil, err
	}

	// Add deployment variables
	if request.Variables != nil {
		for _, variable := range *request.Variables {
			variable.DeploymentID = deployment.ID
			if err := conn.WithContext(ctx).Create(&variable).Error; err != nil {
				// Rollback the deployment creation on failure
				conn.WithContext(ctx).Delete(&deployment)
				return nil, err
			}
		}
	}

	// Return the created deployment
	return &deployment, nil
}
