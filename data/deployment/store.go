package deployment

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
	//err := db.AutoMigrate(&Deployment{})
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Error migrating schema for Deployment model")
	//	os.Exit(1)
	//}

	// Return the store structure
	return &Store{
		db: db,
	}
}

// Create creates a new deployment entry in the database
func (s *Store) Create(entry *Deployment) (*Deployment, *domainErrors.AppError) {
	if err := s.db.Create(entry).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return entry, nil
}

// ReadFirst retrieves a deployment entry from the database
func (s *Store) ReadFirst(id int) (*Deployment, *domainErrors.AppError) {
	entry := &Deployment{}
	if err := s.db.First(entry, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.NotFound)
		return entry, appErr
	}
	return entry, nil
}

// List retrieves all deployment entries from the database with filtering
func (s *Store) List(ctx *gin.Context) ([]Deployment, *domainErrors.AppError) {
	var entries []Deployment
	err := s.db.Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search),
	).Find(&entries).Error

	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return entries, nil
}

// Delete deletes a deployment entry from the database
func (s *Store) Delete(id int) *domainErrors.AppError {
	if err := s.db.Delete(&Deployment{}, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return appErr
	}
	return nil
}

// SetStatus updates the status of a deployment entry in the database
func (s *Store) SetStatus(id int, status DeploymentStatus) *domainErrors.AppError {
	if err := s.db.Model(&Deployment{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return appErr
	}
	return nil
}

// Exists checks if a deployment entry exists in the database
func (s *Store) Exists(id int) bool {
	var count int64
	s.db.Model(&Deployment{}).Where("id = ?", id).Count(&count)
	return count > 0
}
