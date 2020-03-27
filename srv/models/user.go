package models

// User model
type User struct {
	ID        string `json:"id"`
	AuthID    string `json:"auth_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
