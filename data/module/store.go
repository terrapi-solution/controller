package module

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

// Create creates a new module entry in the database
func (s *Store) Create(request Module, ctx context.Context) (Module, error) {
	// Create the user in the database
	if err := s.db.WithContext(ctx).Create(&request).Error; err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return Module{}, domainErrors.NewConflict(err, fmt.Sprintf("Module already exists with the name: %s", request.Name), "ModuleStore.Create")
		}
		return Module{}, domainErrors.NewInternal(err, "Error executing SQL query", "ModuleStore.Create")
	}
	return request, nil
}

// Read retrieves a user module from the database
func (s *Store) Read(id int) (Module, error) {
	var entry Module
	if err := s.db.First(&entry, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Module{}, domainErrors.NewNotFound(err, fmt.Sprintf("Unable to find module with the Id: %d", id), "ModuleStore.Read")
		}
		return Module{}, domainErrors.NewInternal(err, "Error executing SQL query", "ModuleStore.Read")
	}
	return entry, nil
}

// Delete deletes a module entry from the database
func (s *Store) Delete(id int, ctx context.Context) error {
	if err := s.db.WithContext(ctx).Delete(&Module{}, id).Error; err != nil {
		return domainErrors.NewInternal(err, "Error executing SQL query", "ModuleStore.Delete")
	}
	return nil
}

// List retrieves all module entries from the database
func (s *Store) List() ([]Module, error) {
	var entries []Module
	if err := s.db.Find(&entries).Error; err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.List")
	}
	return entries, nil
}

// Exists if a user exists in the database
func (s *Store) Exists(id int) bool {
	var count int64
	if err := s.db.Model(&Module{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// Update updates a module entry in the database
func (s *Store) Update(id int, request Module, ctx context.Context) (Module, error) {
	if err := s.db.WithContext(ctx).Model(&Module{}).Where("id = ?", id).Updates(&request).Error; err != nil {
		return Module{}, domainErrors.NewInternal(err, "Error executing SQL query", "ModuleStore.Update")
	}
	return request, nil
}

// ExistsByName if a module exists in the database
func (s *Store) ExistsByName(name string) bool {
	var count int64
	if err := s.db.Model(&Module{}).Where("lower(name) = lower(?)", name).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
