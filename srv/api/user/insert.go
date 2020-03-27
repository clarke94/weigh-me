package user

import (
	"encoding/json"
	"net/http"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/clarke94/weigh-me/srv/models"
	"github.com/go-chi/render"
)

// InsertUser creates a new user
func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to decode the request body", http.StatusBadRequest)
		return
	}

	insertID, err := driver.InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, insertID)
}
