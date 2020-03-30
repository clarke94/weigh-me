package weight

import (
	"encoding/json"
	"net/http"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/clarke94/weigh-me/srv/models"
	"github.com/go-chi/render"
)

// CreateWeight will create one weight data for a given user
func CreateWeight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var weight models.Weight

	// decode the json request to weight
	err := json.NewDecoder(r.Body).Decode(&weight)
	if err != nil {
		http.Error(w, "Unable to decode the request body", http.StatusBadRequest)
		return
	}

	weight, err = driver.InsertWeight(weight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res models.InsertOneWeightResponse
	res.Status = 200
	res.StatusText = "ok"
	res.URL = r.URL.Path
	res.Ok = true
	res.Name = "Successful"
	res.Message = "One weight successfully created"
	res.Results = weight
	res.Error = ""
	render.JSON(w, r, res)
}
