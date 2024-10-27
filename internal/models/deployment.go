package models

import "time"

type Deployment struct {
	ID         uint                  `gorm:"primaryKey" json:"id"`
	ModuleId   uint                  `json:"-"`
	Name       string                `gorm:"not null" json:"name"`
	Status     DeploymentStatus      `gorm:"not null" json:"status"`
	Module     Module                `gorm:"foreignKey:ModuleId;references:ID" json:"-"`
	Variables  *[]DeploymentVariable `gorm:"foreignKey:DeploymentID;references:ID" json:"-"`
	Activities *[]Activity           `gorm:"foreignKey:DeploymentID;references:ID" json:"-"`
	CreatedAt  time.Time             `gorm:"autoCreateTime" json:"created_at"`
}
