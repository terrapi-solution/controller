package module

import (
	"github.com/terrapi-solution/controller/data"
	"gorm.io/gorm"
)

// Store struct manages interactions with authors store
type Store struct {
	data.Generic[Module]
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// Return the store structure
	return &Store{
		Generic: data.New[Module](db),
	}
}
