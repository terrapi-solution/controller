package models

type DeploymentVariable struct {
	DeploymentID uint   `json:"-"`
	Name         string `gorm:"not null" json:"name"`
	Value        string `gorm:"not null" json:"value"`
}
