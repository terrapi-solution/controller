package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
)

type Generic[T any] struct {
	db *gorm.DB
}

// New creates a new Store struct
func New[T any](db *gorm.DB) Generic[T] {
	return Generic[T]{
		db: db,
	}
}

// Read retrieves an entry from the database
func (s *Generic[T]) Read(id int) (T, error) {
	var entry T
	err := s.db.First(&entry, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entry, domainErrors.NewNotFound(err, fmt.Sprintf("Unable to find entry with the Id: %d", id), "DataStore.Read")
		}
		return entry, domainErrors.NewInternal(err, "Error executing SQL query", "DataStore.Read")
	}
	return entry, nil
}

// Update updates an entry in the database
func (s *Generic[T]) Update(id int, request T, ctx context.Context) (T, error) {
	var entry T
	err := s.db.WithContext(ctx).Model(&entry).Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return entry, domainErrors.NewInternal(err, "Error executing SQL query", "DataStore.Update")
	}
	return request, nil
}

// ExistsByName checks if an entry exists in the database
func (s *Generic[T]) ExistsByName(name string) bool {
	var entry T
	var count int64
	err := s.db.Model(&entry).Where("lower(name) = lower(?)", name).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

// Create creates a new entry in the database
func (s *Generic[T]) Create(request T, ctx context.Context) (T, error) {
	var entry T
	err := s.db.WithContext(ctx).Create(&request).Error
	if err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return entry, domainErrors.NewConflict(err, "Entry already exists", "DataStore.Create")
		}
		return entry, domainErrors.NewInternal(err, "Error executing SQL query", "DataStore.Create")
	}
	return request, nil
}

// Exists checks if an entry exists in the database
func (s *Generic[T]) Exists(id int) bool {
	var entry T
	var count int64
	if err := s.db.Model(&entry).Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// List retrieves all entries from the database
func (s *Generic[T]) List() ([]T, error) {
	var entries []T
	if err := s.db.Find(&entries).Error; err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "DataStore.List")
	}
	return entries, nil
}

// Delete deletes an entry from the database
func (s *Generic[T]) Delete(id int, ctx context.Context) error {
	var entry T
	if err := s.db.WithContext(ctx).Delete(&entry, id).Error; err != nil {
		return domainErrors.NewInternal(err, "Error executing SQL query", "DataStore.Delete")
	}
	return nil
}
