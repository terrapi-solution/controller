package models

type ModuleRequest struct {
	Name   string       `json:"name"`
	Source ModuleSource `json:"source"`
}
