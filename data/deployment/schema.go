package deployment

import (
	"github.com/terrapi-solution/controller/data/activity"
	"github.com/terrapi-solution/controller/data/module"

	"time"
)

type Status string

const (
	Unknown   Status = "unknown"
	Pending   Status = "pending"
	Running   Status = "running"
	Failed    Status = "failed"
	Succeeded Status = "succeeded"
)

type Deployment struct {
	ID         int                  `gorm:"primaryKey"`
	OrderID    int                  `gorm:"not null"`
	ModuleID   int                  `gorm:"not null"`
	Name       string               `gorm:"not null"`
	Status     Status               `gorm:"not null"`
	Module     module.Module        `gorm:"foreignKey:ModuleID;references:ID"`
	Activities *[]activity.Activity `gorm:"foreignKey:DeploymentID;references:ID"`
	CreatedAt  time.Time            `gorm:"autoCreateTime"`
}
