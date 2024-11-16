package module

import (
	"github.com/terrapi-solution/controller/internal/database/trackable"
	"gorm.io/gorm"
	"time"
)

type TypeModule string

const (
	UndefinedType TypeModule = "undefined"
	GitType       TypeModule = "git"
)

type Module struct {
	ID                  int        `gorm:"primaryKey" json:"id" filter:"filterable"`
	Name                string     `gorm:"uniqueIndex;not null" filter:"searchable;filterable"`
	Type                TypeModule `gorm:"not null" json:"type" filter:"filterable"`
	Config              []byte     `gorm:"not null"`
	trackable.CreatedBy `gorm:"not null"`
	trackable.UpdatedBy
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *Module) BeforeCreate(tx *gorm.DB) (err error) {
	return u.CreatedBy.BeforeCreate(tx)
}

func (u *Module) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.UpdatedBy.BeforeUpdate(tx)
}
