package order

import (
	"github.com/terrapi-solution/controller/data/deployment"
	"github.com/terrapi-solution/controller/data/module"
	"github.com/terrapi-solution/controller/data/orderVariable"
	"time"
)

type Type string

const (
	FireAndForget Type = "fire_and_forget"
	Schedule      Type = "schedule"
)

type Order struct {
	ID          int                            `gorm:"primaryKey"`
	Name        string                         `gorm:"index:idx_name,unique"`
	Type        Type                           `gorm:"default:'fire_and_forget'"`
	Schedule    string                         `gorm:"default:'none'"`
	Status      bool                           `gorm:"not null"`
	ModuleID    int                            `gorm:"not null"`
	Module      module.Module                  `gorm:"foreignKey:ModuleID;references:ID"`
	Deployments *[]deployment.Deployment       `gorm:"foreignKey:OrderID;references:ID"`
	Variables   *[]orderVariable.OrderVariable `gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt   time.Time                      `gorm:"autoCreateTime"`
}
