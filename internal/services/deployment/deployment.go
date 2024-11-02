package deployment

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/models"
	"github.com/terrapi-solution/controller/internal/services/module"
	"gorm.io/gorm"
)

type ServiceDeployment struct {
	conn *gorm.DB
}

func NewDeploymentService() *ServiceDeployment {
	return &ServiceDeployment{conn: database.GetInstance()}
}

func (receiver *ServiceDeployment) Create(ctx context.Context, request models.DeploymentRequest) (*models.Deployment, error) {
	// Check if the module exists
	moduleService := module.NewModuleService()
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
	if err := receiver.conn.WithContext(ctx).Create(&deployment).Error; err != nil {
		return nil, err
	}

	// Add deployment variables
	if request.Variables != nil {
		for _, variable := range *request.Variables {
			variable.DeploymentID = deployment.ID
			if err := receiver.conn.WithContext(ctx).Create(&variable).Error; err != nil {
				// Rollback the deployment creation on failure
				receiver.conn.WithContext(ctx).Delete(&deployment)
				return nil, err
			}
		}
	}

	// Return the created deployment
	return &deployment, nil
}
