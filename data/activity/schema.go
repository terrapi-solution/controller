package activity

import "time"

type Activity struct {
	ID           int       `gorm:"primaryKey"`
	DeploymentID int       `gorm:"not null"`
	Pointer      string    `gorm:"not null" `
	Message      string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
