package service

import (
	"context"
	"errors"
	"github.com/terrapi-solution/controller/internal/database"
	model "github.com/terrapi-solution/controller/internal/models"
)

type Deployment struct {
}

func NewDeploymentService() *Deployment {
	return &Deployment{}
}

func (a *Deployment) List(ctx context.Context, page, pageSize int) ([]model.Deployment, error) {
	// Define the list of activities
	var entities []model.Deployment

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

func (a *Deployment) Get(ctx context.Context, deploymentId int) (*model.Deployment, error) {
	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return nil, errors.New("database instance is not initialized")
	}

	// Get the deployment
	var entity *model.Deployment
	if err := conn.WithContext(ctx).
		Where("id = ?", deploymentId).
		First(&entity).Error; err != nil {
		return nil, err
	}

	// Return the deployment
	return entity, nil
}
