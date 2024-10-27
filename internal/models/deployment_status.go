package models

type DeploymentStatus string

const (
	Unknown   DeploymentStatus = "unknown"
	Pending   DeploymentStatus = "pending"
	Running   DeploymentStatus = "running"
	Failed    DeploymentStatus = "failed"
	Succeeded DeploymentStatus = "succeeded"
)
