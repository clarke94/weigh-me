package user

import (
	"net/http"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/clarke94/weigh-me/srv/models"
	"github.com/go-chi/render"
)

// GetUserByAuthID will return a single user from the ID provided by the Auth API
func GetUserByAuthID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/x-www-form-urlencoded")
	// get the userid from the request params, key is "id"
	id := r.URL.Query().Get("id")

	// call the getUser function with user id to retrieve a single user
	user, err := driver.GetUserByAuthID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res models.GetOneUserResponse
	res.Status = 200
	res.StatusText = "ok"
	res.URL = r.URL.Path
	res.Ok = true
	res.Name = "Successful"
	res.Message = "User was successfully returned"
	res.Results = user
	res.Error = ""
	render.JSON(w, r, res)
}
