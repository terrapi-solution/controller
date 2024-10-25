package database

import (
	"github.com/terrapi-solution/controller/internal/models/database"
)

func (s *DatabaseConnection) Migrator() error {
	// Define all database models
	models := []interface{}{
		&database.Module{},
		&database.ModuleSource{},
		&database.Deployment{},
		&database.DeploymentVariable{},
	}

	// Execute auto migration for all models
	for _, model := range models {
		if err := s.Conn.AutoMigrate(model); err != nil {
			return err
		}
	}

	// return nil if no error
	return nil
}
