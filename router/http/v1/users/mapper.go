package users

import (
	"github.com/terrapi-solution/controller/data/user"
)

func toUserResponse(u user.User) *UserResponse {
	return &UserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		Status:     u.Status,
		LastActive: u.LastActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
