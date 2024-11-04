package user

import "time"

type User struct {
	ID             int    `gorm:"primaryKey"`
	Email          string `gorm:"not null;unique"`
	HashedPassword string
	FirstName      string
	LastName       string
	IsActive       bool
	LastActive     time.Time
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
