package core

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/controller/internal/database"
	"sync"
)

type Core struct {
	Config *config.Config
}

var (
	instance *Core
	once     sync.Once
)

func GetInstance() *Core {
	// Make sure the instance is created only once.
	once.Do(func() {
		log.Info().Msg("Creating a new core service instance")
		instance = &Core{}

		log.Info().Msg("Initializing the core service")
		instance.initializeConfiguration()
		instance.initializeDatabase()

		log.Info().Msg("Core service initialized")
	})
	return instance
}

func (c *Core) initializeConfiguration() {
	log.Info().Msg("Initializing configuration")
	c.Config = config.GetInstance()
}

func (c *Core) initializeDatabase() {
	log.Info().Msg("Initializing database service")
	database.Initialize(c.Config.Datastore)
	database.CreateModel()
}
