package activity

import (
	"github.com/gin-gonic/gin"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
)

// Store struct manages interactions with authors store
type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// Migrate the schema
	//err := db.AutoMigrate(&Activity{})
	//if err != nil {
	//		log.Fatal().Err(err).Msg("Error migrating schema for Activity model")
	//		os.Exit(1)
	//	}

	// Return the store structure
	return &Store{
		db: db,
	}
}

// Create creates a new activity entry in the database
func (s *Store) Create(entry *Activity) (*Activity, *domainErrors.AppError) {
	if err := s.db.Create(entry).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return entry, nil
}

// Read retrieves an activity entry from the database
func (s *Store) Read(id uint) (*Activity, *domainErrors.AppError) {
	entry := &Activity{}
	if err := s.db.First(entry, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.NotFound)
		return nil, appErr
	}
	return entry, nil
}

// List retrieves all activity entries from the database with filtering
func (s *Store) List(ctx *gin.Context) (*[]Activity, *domainErrors.AppError) {
	var entries []Activity
	err := s.db.Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search),
	).Find(&entries).Error

	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return &entries, nil
}

// ListByDeploymentID retrieves all activity entries from the database for a specific deployment
func (s *Store) ListByDeploymentID(ctx *gin.Context, deploymentID int) ([]Activity, *domainErrors.AppError) {
	var entries []Activity
	err := s.db.Scopes(
		filter.FilterByQuery(ctx, filter.Paginate),
	).Where("deployment_id = ?", deploymentID).Find(&entries).Error

	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return entries, nil
}
