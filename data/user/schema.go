package user

import "time"

type User struct {
	ID         int       `gorm:"primaryKey" filter:"filterable"`
	Subject    string    `gorm:"not null;unique"`
	UserName   string    `gorm:"not null;unique" filter:"searchable;filterable"`
	FirstName  string    `filter:"filterable" filter:"searchable;filterable"`
	LastName   string    `filter:"filterable" filter:"searchable;filterable"`
	Email      string    `gorm:"not null" filter:"searchable;filterable"`
	Status     bool      `filter:"filterable"`
	LastActive time.Time `filter:"filterable"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
