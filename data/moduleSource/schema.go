package moduleSource

type ModuleSource struct {
	ModuleID   int    `gorm:"not null"`
	Repository string `gorm:"not null"`
	Branch     string `gorm:"not null"`
	Path       string `gorm:"not null"`
}
