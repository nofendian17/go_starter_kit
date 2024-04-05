package response

type HealthResponse struct {
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
	CPU     string `json:"cpu"`
	Memory  string `json:"memory"`
}
