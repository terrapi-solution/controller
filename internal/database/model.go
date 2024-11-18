package database

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/plan"
	"github.com/terrapi-solution/controller/data/planVariable"
)

func CreateModel() {
	// Define all database models
	entities := []interface{}{
		&module.Module{},
		&plan.Plan{},
		&planVariable.PlanVariable{},
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
