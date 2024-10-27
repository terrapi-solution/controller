package models

type ModuleSource struct {
	ModuleID   uint   `json:"-"`
	Repository string `gorm:"not null" json:"repository"`
	Branch     string `gorm:"not null" json:"branch"`
	Path       string `gorm:"not null" json:"path"`
}
