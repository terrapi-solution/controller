package models

import "time"

type Activity struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DeploymentID uint      `gorm:"not null" json:"-"`
	Pointer      string    `gorm:"not null" json:"pointer"`
	Message      string    `gorm:"not null" json:"message"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
