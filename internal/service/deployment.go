package service

import "github.com/terrapi-solution/controller/internal/database"

type Deployment struct {
	DB *database.DatabaseConnection
}

func NewDeploymentService(connection *database.DatabaseConnection) *Deployment {
	return &Deployment{DB: connection}
}
