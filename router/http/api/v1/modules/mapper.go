package modules

import (
	data "github.com/terrapi-solution/controller/data/module"
	domain "github.com/terrapi-solution/controller/domain/module"
)

func toModuleDto(entity data.Module) ModuleResponseDto {
	return ModuleResponseDto{
		ID:        entity.ID,
		Name:      entity.Name,
		Type:      entity.Type,
		CreatedBy: entity.CreatedBy.CreatedBy,
		UpdatedBy: entity.UpdatedBy.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func toGitConfigDto(entity domain.GitConfigRequest) GitConfigDto {
	return GitConfigDto{
		Repository: entity.Repository,
		Branch:     entity.Branch,
		Username:   entity.Username,
	}
}
