package module

import (
	"github.com/terrapi-solution/controller/data/moduleSource"
	"time"
)

type moduleType string

const (
	Git moduleType = "git"
)

type Module struct {
	ID        int                       `gorm:"primaryKey" json:"id" filter:"filterable"`
	Name      string                    `gorm:"uniqueIndex;not null" filter:"searchable;filterable"`
	Type      moduleType                `gorm:"not null" json:"type" filter:"filterable"`
	Source    moduleSource.ModuleSource `gorm:"foreignKey:ModuleID;references:ID"`
	CreatedAt time.Time                 `gorm:"autoCreateTime"`
	UpdatedAt time.Time                 `gorm:"autoUpdateTime"`
}
