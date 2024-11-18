package module

import (
	"github.com/terrapi-solution/controller/internal/database/trackable"
	"gorm.io/gorm"
	"time"
)

type TypeModule string

const (
	// UndefinedType represents an undefined module type
	UndefinedType TypeModule = "undefined"
	// GitType represents a git module type
	GitType TypeModule = "git"
)

type Module struct {
	// ID defines the unique identifier.
	ID int `gorm:"primaryKey" json:"id" filter:"filterable"`

	// Name defines the unique name of the module.
	Name string `gorm:"uniqueIndex;not null" filter:"searchable;filterable"`

	// Type defines the type of the module.
	Type TypeModule `gorm:"not null" json:"type" filter:"filterable"`

	// Config defines the configuration of the module.
	// It is a JSON object that contains the configuration of the module.
	Config []byte

	// Audit fields
	trackable.CreatedBy
	CreatedAt time.Time `gorm:"autoCreateTime"`
	trackable.UpdatedBy
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *Module) BeforeCreate(tx *gorm.DB) (err error) {
	return u.CreatedBy.BeforeCreate(tx)
}

func (u *Module) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.UpdatedBy.BeforeUpdate(tx)
}
