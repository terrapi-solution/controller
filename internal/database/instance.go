package database

import (
	"fmt"
	"github.com/terrapi-solution/controller/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DatabaseConnection struct {
	Conn *gorm.DB
}

var (
	dbInstance *DatabaseConnection
	once       sync.Once
)

func GetDatabaseConnection(config config.Datastore) *DatabaseConnection {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Host, config.Username, config.Password, config.Database, config.Port,
		)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		dbInstance = &DatabaseConnection{Conn: conn}
		fmt.Println("Database connection established")
	})

	return dbInstance
}
