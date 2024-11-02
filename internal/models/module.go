package models

import (
	"time"
)

type Module struct {
	ID        uint         `gorm:"primaryKey" json:"id" filter:"filterable"`
	Name      string       `gorm:"uniqueIndex;not null" json:"name" filter:"searchable;filterable"`
	Type      ModuleType   `gorm:"not null" json:"type" filter:"filterable"`
	Source    ModuleSource `gorm:"foreignKey:ModuleID;references:ID" json:"-"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
