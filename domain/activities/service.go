package activities

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/data/activity"
	"github.com/terrapi-solution/controller/data/deployment"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	activity   *activity.Store
	deployment *deployment.Store
}

// New creates a new Service struct
func New(db *gorm.DB) *Service {
	return &Service{
		activity:   activity.New(db),
		deployment: deployment.New(db),
	}
}

// Create creates a new activity entry in the database
func (s *Service) Create(request ActivityRequest) (*activity.Activity, *domainErrors.AppError) {
	// Check if the deployment is existing
	// TODO: Add cache for deployment to avoid multiple database calls
	dpl, err := s.deployment.ReadFirst(request.DeploymentID)
	if err != nil {
		if err.Type == domainErrors.NotFound {
			log.Debug().Err(err).Msgf("Deployment not found for ID: %d", request.DeploymentID)
			return nil, domainErrors.NewAppError(nil, domainErrors.ValidationError)
		} else {
			log.Debug().Err(err).Msgf("Error reading deployment for ID: %d", request.DeploymentID)
			return nil, err
		}
	}

	// Check if the deployment is active
	if dpl.Status != "running" {
		log.Debug().Msgf("Deployment is not running for ID: %d", request.DeploymentID)
		return nil, domainErrors.NewAppError(nil, domainErrors.NotAuthorized)
	}

	// Create the activity to the database
	entry, appErr := s.activity.Create(toActivityDBModel(request))
	if appErr != nil {
		log.Debug().Err(appErr).Msg("Error creating new activity")
		return nil, appErr
	}
	return entry, nil
}

// Read retrieves an activity entry from the database
func (s *Service) Read(id uint) (*activity.Activity, *domainErrors.AppError) {
	entry, err := s.activity.Read(id)
	if err != nil {
		if err.Type == domainErrors.NotFound {
			log.Debug().Err(err).Msgf("Activity not found for ID: %d", id)
			return nil, err
		} else {
			log.Debug().Err(err).Msgf("Error reading activity for ID: %d", id)
			return nil, err
		}
	}
	return entry, nil
}
