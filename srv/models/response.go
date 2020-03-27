package models

// WeightResponse model for http responses
type WeightResponse struct {
	Status     int      `json:"status"`
	StatusText string   `json:"statusText"`
	URL        string   `json:"url"`
	Ok         bool     `json:"ok"`
	Name       string   `json:"name"`
	Message    string   `json:"message"`
	Results    []Weight `json:"results"`
	Error      string   `json:"error"`
}

// UserResponse model for http responses
type UserResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	URL        string `json:"url"`
	Ok         bool   `json:"ok"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}
