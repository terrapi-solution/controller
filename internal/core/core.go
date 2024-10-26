package core

import (
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/config"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/service"
	"sync"
)

type Core struct {
	Config      *config.Config
	DB          *database.DatabaseConnection
	Deployment  *service.Deployment
	Activity    *service.Activity
	HealthCheck *service.HealthService
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
		instance.initializeDeployment()
		instance.initializeActivity()

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
	c.DB = database.GetDatabaseConnection(c.Config.Datastore)
	c.DB.CreateModel()
}

func (c *Core) initializeDeployment() {
	log.Info().Msg("Initializing deployment service")
}

func (c *Core) initializeActivity() {
	log.Info().Msg("Initializing activity service")
}

func (c *Core) Dispose() {
	log.Info().Msg("Disposing the core service")
	// Dispose the services
	c.Activity = nil
	c.Deployment = nil

	// Close the database connection
	c.DB = nil

	// Dispose the configuration
	c.Config = nil
}
