package module

// ModuleRequest struct represents the module configuration
type ModuleRequest struct {
	Name string `json:"name" validate:"required,min=4,max=20"`
}

// GitConfigRequest struct represents the git configuration for a module
type GitConfigRequest struct {
	Repository string `json:"repository" validate:"required"`
	Branch     string `json:"branch" validate:"required,min=4,max=20"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
