package database

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/models/database"
)

func (s *DatabaseConnection) CreateModel() error {
	// Define all database models
	models := []interface{}{
		&database.Module{},
		&database.ModuleSource{},
		&database.Deployment{},
		//&database.DeploymentVariable{},
		&database.Activity{},
	}

	// Execute auto migration for all models
	for _, model := range models {
		if err := s.Conn.AutoMigrate(model); err != nil {
			return err
		}
	}

	// return nil if no error
	log.Info().Msg("Database models created")
	return nil
}
