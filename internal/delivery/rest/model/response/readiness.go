package response

type ReadinessResponse struct {
	Status string  `json:"status"`
	Checks []Check `json:"checks"`
}
