package user

// User Represents information of current user
type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}
