package database

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/data/activity"
	"github.com/terrapi-solution/controller/data/deployment"
	"github.com/terrapi-solution/controller/data/deploymentVariable"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/moduleSource"
)

func CreateModel() {
	// Define all database models
	entities := []interface{}{
		&module.Module{},
		&moduleSource.ModuleSource{},
		&deployment.Deployment{},
		&deploymentVariable.DeploymentVariable{},
		&activity.Activity{},
	}

	// Execute auto migration for all models
	for _, entity := range entities {
		if err := instance.AutoMigrate(entity); err != nil {
			log.Error().Err(err).
				Msg("Failed to create database model")
		}
	}

	log.Info().Msg("Database models created")
}
