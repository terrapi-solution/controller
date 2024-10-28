package models

type ModuleRequest struct {
	Name   string       `json:"name"`
	Type   ModuleType   `gorm:"not null" json:"type"`
	Source ModuleSource `json:"source"`
}
