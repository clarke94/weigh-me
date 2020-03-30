package models

// Weight model
type Weight struct {
	ID     string  `json:"id"`
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
}
