package deployments

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/data/activity"
	"github.com/terrapi-solution/controller/data/deployment"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/orderVariable"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	module     *module.Store
	deployment *deployment.Store
	variable   *orderVariable.Store
	Activity   *activity.Store
}

// New creates a new Service struct
func New(db *gorm.DB) Service {
	return Service{
		module:     module.New(db),
		deployment: deployment.New(db),
		variable:   orderVariable.New(db),
		Activity:   activity.New(db),
	}
}

// Create creates a new deployment entry in the database
func (s *Service) Create(request DeploymentRequest) (*deployment.Deployment, *domainErrors.AppError) {
	// Check if the module is existing
	// TODO: Add cache for module to avoid multiple database calls
	_, err := s.module.Read(request.ModuleID)
	if err != nil {
		if err.Type == domainErrors.NotFound {
			log.Debug().Err(err).Msgf("Module not found for ID: %d", request.ModuleID)
			return nil, domainErrors.NewAppError(nil, domainErrors.ValidationError)
		} else {
			log.Debug().Err(err).Msgf("Error reading module for ID: %d", request.ModuleID)
			return nil, err
		}
	}

	// Create the activity to the database
	entry, dplErr := s.deployment.Create(toDeploymentDBModel(request))
	if dplErr != nil {
		log.Debug().Err(dplErr).Msg("Error creating new deployment")
		return nil, domainErrors.NewAppError(dplErr, domainErrors.RepositoryError)
	}

	// Create the deployment variables to the database
	if request.Variables != nil {
		for _, variable := range request.Variables {
			if _, varErr := s.variable.Create(toDeploymentVariableDBModel(variable, entry.ID)); varErr != nil {
				log.Debug().Err(varErr).Msg("Error creating new deployment variables")
				if delErr := s.deployment.Delete(entry.ID); delErr != nil {
					log.Debug().Err(delErr).Msgf("Error deleting deployment with ID: %d", entry.ID)
					return nil, domainErrors.NewAppError(delErr, domainErrors.UnknownError)
				}
				return nil, domainErrors.NewAppError(varErr, domainErrors.UnknownError)
			}
		}
	}

	// Set the deployment status to pending
	if staErr := s.deployment.SetStatus(entry.ID, deployment.Pending); staErr != nil {
		log.Debug().Err(staErr).Msgf("Error deleting deployment with ID: %d", entry.ID)
	}

	return entry, nil
}

// ReadFirst reads a deployment entry from the database
func (s *Service) ReadFirst(id int) (*deployment.Deployment, *domainErrors.AppError) {
	data, err := s.deployment.ReadFirst(id)
	return data, err
}

// List reads all deployment entries from the database
func (s *Service) List(ctx *gin.Context) ([]deployment.Deployment, *domainErrors.AppError) {
	return s.deployment.List(ctx)
}

// Delete deletes a deployment entry from the database
func (s *Service) Delete(id int) error {
	if s.deployment.Exists(id) == false {
		return domainErrors.NewAppError(nil, domainErrors.NotFound)
	}

	if err := s.deployment.Delete(id); err != nil {
		log.Debug().Err(err).Msgf("Error deleting deployment with ID: %d", id)
		return domainErrors.NewAppError(err, domainErrors.UnknownError)
	}

	return nil
}

func (s *Service) ListActivities(id int, ctx *gin.Context) ([]activity.Activity, *domainErrors.AppError) {
	// Check if the deployment exists
	if s.deployment.Exists(id) == false {
		return nil, domainErrors.NewAppError(nil, domainErrors.NotFound)
	}

	// Get the activities from the database
	return s.Activity.ListByDeploymentID(ctx, id)
}
