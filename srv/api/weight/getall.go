package weight

import (
	"net/http"
	"strconv"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
)

// GetAllWeights will return all weight values in database
func GetAllWeights(w http.ResponseWriter, r *http.Request) {
	weights, err := driver.GetAllWeights()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, weights)
}

// GetAllWeightsByUserID will return all weight values in database for a given user ID
func GetAllWeightsByUserID(w http.ResponseWriter, r *http.Request) {
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

	render.JSON(w, r, weights)
}
