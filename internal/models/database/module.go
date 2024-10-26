package database

import (
	"time"
)

type Module struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `gorm:"uniqueIndex;not null"`
	Source    ModuleSource `gorm:"foreignKey:ModuleID;references:ID"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
}

type ModuleSource struct {
	ModuleID   uint
	Repository string `gorm:"not null"`
	Branch     string `gorm:"not null"`
	Path       string `gorm:"not null"`
}
