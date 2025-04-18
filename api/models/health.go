package models

type HealthStatus struct {
	Status       string            `json:"status"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
	Error        string            `json:"error,omitempty"`
}
