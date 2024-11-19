package plan

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
)

// Store struct manages interactions with authors store
type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// Return the store structure
	return &Store{
		db: db,
	}
}

// ExistsByName if a plan exists in the database
func (s *Store) ExistsByName(name string) bool {
	var count int64
	if err := s.db.Model(&Plan{}).Where("lower(name) = lower(?)", name).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (s *Store) Create(request Plan, ctx context.Context) (Plan, error) {
	// Create the user in the database
	if err := s.db.WithContext(ctx).Create(&request).Error; err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return Plan{}, domainErrors.NewConflict(err, fmt.Sprintf("Plan already exists with the name: %s", request.Name), "PlanStore.Create")
		}
		return Plan{}, domainErrors.NewInternal(err, "Error executing SQL query", "PlanStore.Create")
	}
	return request, nil
}

// Read retrieves a user module from the database
func (s *Store) Read(id int) (Plan, error) {
	var entry Plan
	if err := s.db.First(&entry, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Plan{}, domainErrors.NewNotFound(err, fmt.Sprintf("Unable to find plan with the Id: %d", id), "PlanStore.Read")
		}
		return Plan{}, domainErrors.NewInternal(err, "Error executing SQL query", "PlanStore.Read")
	}
	return entry, nil
}

// Update updates a module entry in the database
func (s *Store) Update(id int, request Plan, ctx context.Context) (Plan, error) {
	if err := s.db.WithContext(ctx).Model(&Plan{}).Where("id = ?", id).Updates(&request).Error; err != nil {
		return Plan{}, domainErrors.NewInternal(err, "Error executing SQL query", "PlanStore.Update")
	}
	return request, nil
}
