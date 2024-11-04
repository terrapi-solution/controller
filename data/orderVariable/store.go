package orderVariable

import (
	"github.com/terrapi-solution/controller/data/deploymentVariable"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
)

// Store struct manages interactions with authors store
type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// Migrate the schema
	//err := db.AutoMigrate(&DeploymentVariable{})
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Error migrating schema for DeploymentVariable model")
	//	os.Exit(1)
	//}

	// Return the store structure
	return &Store{
		db: db,
	}
}

// Create creates a new deployment variable entry in the database
func (s *Store) Create(entry *deploymentVariable.DeploymentVariable) (*deploymentVariable.DeploymentVariable, *domainErrors.AppError) {
	if err := s.db.Create(entry).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return entry, nil
}

// Read retrieves a deployment variable entry from the database
func (s *Store) Read(id uint) (*deploymentVariable.DeploymentVariable, *domainErrors.AppError) {
	entry := &deploymentVariable.DeploymentVariable{}
	if err := s.db.First(entry, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.NotFound)
		return nil, appErr
	}
	return entry, nil
}
