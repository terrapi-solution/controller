package plan

import (
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/planVariable"
	"github.com/terrapi-solution/controller/internal/database/trackable"
	"gorm.io/gorm"
	"time"
)

// State represents the state of a plan
type State string

// Type represents the type of the plan
type Type string

const (
	// PendingState represents the initial status of a plan once it has been created.
	PendingState State = "pending"
	// RunningState represents the status of a plan when it is being executed
	RunningState State = "running"
	// ErroredState represents the status of a plan when it has encountered an error
	ErroredState State = "errored"
	// CanceledState represents the status of a plan when it has been canceled
	CanceledState State = "canceled"
	// FinishedState represents the status of a plan when it has been successfully executed
	FinishedState State = "finished"
	// UnreachableState represents the status of a plan when it has been unreachable
	UnreachableState State = "unreachable"

	// DefaultType represents a plan that is executed once and does not have a schedule
	DefaultType Type = "default"
	// ScheduleType represents a plan that is executed on a schedule
	ScheduleType Type = "schedule"
)

// Plan represents a plan with its details and audit fields.
type Plan struct {
	// ID defines the unique identifier.
	ID int `gorm:"primaryKey"`

	// Name defines the unique name of the plan.
	// It is also used by the state manager to map real world resources to your configuration
	Name string `gorm:"index:idx_name,unique"`

	// Type defines the type of the plan.
	Type Type `gorm:"default:'default'"`

	// States defines the state of the execution.
	State State `gorm:"default:'pending'"`

	// Schedule defines the schedule of the execution.
	// if the plan is of type fire_and_forget, this field is set to "none"
	Schedule string `gorm:"default:'none'"`

	// ModuleID defines the identifier of the module.
	ModuleID int           `gorm:"not null"`
	Module   module.Module `gorm:"foreignKey:ModuleID;references:ID"`

	// Variables defines the variables of the plan.
	Variables []planVariable.PlanVariable `gorm:"foreignKey:PlanID;references:ID"`

	// Audit fields
	trackable.CreatedBy
	CreatedAt time.Time `gorm:"autoCreateTime"`
	trackable.UpdatedBy
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Sanitize applies the business logic to the plan,
// it used to ensure that the plan is in a valid state.
func (u *Plan) Sanitize(tx *gorm.DB) {
	if u.Type == DefaultType && u.Schedule != "none" {
		tx.Statement.SetColumn("schedule", "none")
	}
}

// BeforeCreate executes the logic before creating a plan.
func (u *Plan) BeforeCreate(tx *gorm.DB) (err error) {
	u.Sanitize(tx)
	return u.CreatedBy.BeforeCreate(tx)
}

// BeforeUpdate executes the logic before updating a plan.
func (u *Plan) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Sanitize(tx)
	return u.UpdatedBy.BeforeUpdate(tx)
}
