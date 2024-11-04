package moduleSource

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
	//err := db.AutoMigrate(&ModuleSource{})
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Error migrating schema for ModuleSource model")
	//	os.Exit(1)
	//}

	// Return the store structure
	return &Store{
		db: db,
	}
}

// Create creates a new module source entry in the database
func (s *Store) Create(entry *ModuleSource) (*ModuleSource, *domainErrors.AppError) {
	if err := s.db.Create(entry).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return entry, nil
}

// Read retrieves a module source entry from the database
func (s *Store) Read(id uint) (*ModuleSource, *domainErrors.AppError) {
	entry := &ModuleSource{}
	if err := s.db.First(entry, id).Error; err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.NotFound)
		return nil, appErr
	}
	return entry, nil
}

// List retrieves all module source entries from the database with filtering
func (s *Store) List(ctx *gin.Context) (*[]ModuleSource, *domainErrors.AppError) {
	var entries []ModuleSource
	err := s.db.Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search),
	).Find(&entries).Error

	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return &entries, nil
}
