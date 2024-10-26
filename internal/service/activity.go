package service

import "github.com/terrapi-solution/controller/internal/database"

type Activity struct {
	DB *database.DatabaseConnection
}

func NewActivityService(connection *database.DatabaseConnection) *Activity {
	return &Activity{DB: connection}
}

func (a *Activity) Create() {
}
