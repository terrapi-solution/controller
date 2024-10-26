package database

import "time"

type Deployment struct {
	ID         uint `gorm:"primaryKey"`
	ModuleId   uint
	Name       string                `gorm:"not null"`
	Status     DeploymentStatus      `gorm:"not null"`
	Module     Module                `gorm:"foreignKey:ModuleId;references:ID"`
	Variables  *[]DeploymentVariable `gorm:"foreignKey:DeploymentID;references:ID"`
	Activities *[]Activity           `gorm:"foreignKey:DeploymentID;references:ID"`
	CreatedAt  time.Time
}

type DeploymentVariable struct {
	DeploymentID int
	Name         string
	Value        string
}

type DeploymentStatus string

const (
	Unknown   DeploymentStatus = "unknown"
	Pending   DeploymentStatus = "pending"
	Running   DeploymentStatus = "running"
	Failed    DeploymentStatus = "failed"
	Succeeded DeploymentStatus = "succeeded"
)
