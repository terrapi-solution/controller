package planVariable

import (
	"github.com/terrapi-solution/controller/data"
	"gorm.io/gorm"
)

// Store struct manages interactions with authors store
type Store struct {
	data.Generic[PlanVariable]
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	return &Store{
		Generic: data.New[PlanVariable](db),
	}
}
