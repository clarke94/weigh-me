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

// GetOneUserResponse model for http responses
type GetOneUserResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	URL        string `json:"url"`
	Ok         bool   `json:"ok"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Results    User   `json:"results"`
	Error      string `json:"error"`
}

// InsertOneUserResponse model for http responses
type InsertOneUserResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	URL        string `json:"url"`
	Ok         bool   `json:"ok"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Results    string `json:"results"`
	Error      string `json:"error"`
}
