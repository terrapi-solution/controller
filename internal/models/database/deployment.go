package database

import "time"

type Deployment struct {
	ID        uint `gorm:"primaryKey"`
	ModuleId  uint
	Name      string
	Variables *[]DeploymentVariable `gorm:"foreignKey:DeploymentID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeploymentVariable struct {
	DeploymentID int
	Name         string
	Value        string
}
