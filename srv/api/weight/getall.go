package weight

import (
	"net/http"
	"strconv"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/clarke94/weigh-me/srv/models"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

// GetAllWeights will return all weight values in database
func GetAllWeights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	var res models.WeightResponse

	weights, err := driver.GetAllWeights()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Status = 200
	res.StatusText = "ok"
	res.URL = r.URL.Path
	res.Ok = true
	res.Name = "Successful"
	res.Message = "All weights successfully returned"
	res.Results = weights
	res.Error = ""
	render.JSON(w, r, res)
}

// GetAllWeightsByUserID will return all weight values in database for a given user ID
func GetAllWeightsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	var res models.WeightResponse

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	weights, err := driver.GetAllWeightsByUserID(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Status = 200
	res.StatusText = "ok"
	res.URL = r.URL.Path
	res.Ok = true
	res.Name = "Successful"
	res.Message = "All weights successfully returned"
	res.Results = weights
	res.Error = ""
	render.JSON(w, r, weights)
}
