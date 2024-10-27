package database

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/models"
)

func CreateModel() {
	// Define all database models
	models := []interface{}{
		&models.Module{},
		&models.ModuleSource{},
		&models.Deployment{},
		&models.DeploymentVariable{},
		&models.Activity{},
	}

	// Execute auto migration for all models
	for _, model := range models {
		if err := instance.AutoMigrate(model); err != nil {
			log.Error().Err(err).
				Msg("Failed to create database model")
		}
	}

	log.Info().Msg("Database models created")
}
