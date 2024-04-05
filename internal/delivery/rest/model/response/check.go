package response

type Check struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}
