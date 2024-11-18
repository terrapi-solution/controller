package planVariable

import (
	"github.com/terrapi-solution/controller/internal/database/trackable"
	"gorm.io/gorm"
	"time"
)

// Category defines the category of the variable.
// It used to define the scope of the variable.
type Category string

const (
	RunnerCategory Category = "runner"
	EnvCategory    Category = "env"
)

type PlanVariable struct {
	// ID defines the unique identifier of the plan variable.
	ID int `gorm:"primaryKey" filter:"filterable"`

	// PlanID defines the unique identifier of the plan.
	PlanID int `gorm:"not null" filter:"filterable"`

	// Key defines the unique identifier of the variable.
	Key string `gorm:"not null" filter:"filterable"`

	// Value defines the value of the variable.
	Value string

	// Category define the scope of the variable.
	Category Category `gorm:"not null;default:'runner'" filter:"filterable"`

	// Sensitive defines if the value is sensitive.
	// If true, the variable is written once and not visible thereafter.
	Sensitive bool `gorm:"not null" filter:"filterable"`

	// Audit fields
	trackable.CreatedBy
	CreatedAt time.Time `gorm:"autoCreateTime"`
	trackable.UpdatedBy
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *PlanVariable) BeforeCreate(tx *gorm.DB) (err error) {
	return u.CreatedBy.BeforeCreate(tx)
}

func (u *PlanVariable) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.UpdatedBy.BeforeUpdate(tx)
}
