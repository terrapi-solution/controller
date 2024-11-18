package plan

import "gorm.io/gorm"

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
