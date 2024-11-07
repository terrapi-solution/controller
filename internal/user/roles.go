package user

// Role represents the role of a user
type Role string

const (
	RoleAdmin   Role = "admin"
	RoleMember  Role = "member"
	RoleAuditor Role = "auditor"
)
