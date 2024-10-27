package database

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// GetInstance returns the database instance
func GetInstance() *gorm.DB {
	if instance == nil {
		log.Fatal().Msg("Database instance is not initialized")
		return nil
	} else {
		return instance
	}
}

func Initialize(config config.Datastore) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Host, config.Username, config.Password, config.Database, config.Port,
		)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}

		instance = conn
	})
}
