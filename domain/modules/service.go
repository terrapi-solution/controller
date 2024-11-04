package modules

import (
	"github.com/terrapi-solution/controller/data/module"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	module *module.Store
}

// New creates a new Service struct
func New(db *gorm.DB) *Service {
	return &Service{
		module: module.New(db),
	}
}
