package module

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
	//err := db.AutoMigrate(&Module{})
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Error migrating schema for Module model")
	//	os.Exit(1)
	//}

	// Return the store structure
	return &Store{
		db: db,
	}
}

// Create creates a new module entry in the database
func (s *Store) Create(entry *Module) (*Module, *domainErrors.AppError) {
	if err := s.db.Create(entry).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return entry, nil
}

// Read retrieves a module entry from the database
func (s *Store) Read(id int) (*Module, *domainErrors.AppError) {
	entry := &Module{}
	if err := s.db.First(entry, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.NotFound)
		return nil, appErr
	}
	return entry, nil
}

// List retrieves all module entries from the database with filtering
func (s *Store) List(ctx *gin.Context) (*[]Module, *domainErrors.AppError) {
	var entries []Module
	err := s.db.Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search),
	).Find(&entries).Error

	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return &entries, nil
}
