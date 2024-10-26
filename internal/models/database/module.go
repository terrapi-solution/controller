package database

import (
	"time"
)

type Module struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Source    ModuleSource `gorm:"foreignKey:ModuleID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModuleSource struct {
	ModuleID   uint
	Repository string
	Branch     string
	Path       string
}
