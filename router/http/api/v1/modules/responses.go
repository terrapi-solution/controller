package modules

import (
	"github.com/terrapi-solution/controller/data/module"
	"time"
)

type ModuleResponseDto struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Type      module.TypeModule `json:"type"`
	CreatedBy string            `json:"created_by"`
	UpdatedBy string            `json:"updated_by"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// ListResponseDto struct defines books list response structure
type ListResponseDto struct {
	Data []ModuleResponseDto `json:"data"`
}

// GitConfigDto struct defines the git configuration structure
type GitConfigDto struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
	Username   string `json:"username"`
}
