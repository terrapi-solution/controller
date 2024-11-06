package users

import "time"

// UserResponse struct defines user response structure
type UserResponse struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Status     bool      `json:"status"`
	LastActive time.Time `json:"last_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ListResponse struct defines books list response structure
type ListResponse struct {
	Data []UserResponse `json:"data"`
}
