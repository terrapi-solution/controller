package database

import "time"

type Activity struct {
	ID           uint      `gorm:"primaryKey"`
	DeploymentID uint      `gorm:"not null"`
	Pointer      string    `gorm:"not null"`
	Message      string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
