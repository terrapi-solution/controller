package users

// StatusRequest is used to update a user status.
type StatusRequest struct {
	Status bool `json:"status"`
}
