package activities

import "time"

// ActivityResponse struct defines response fields
type ActivityResponse struct {
	ID           int       `json:"id"`
	DeploymentID int       `json:"deploymentId"`
	Pointer      string    `json:"type"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `gorm:"createdAt"`
}

// ListResponse struct defines books list response structure
type ListResponse struct {
	Data []ActivityResponse `json:"data"`
}
