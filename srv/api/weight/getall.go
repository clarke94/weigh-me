package weight

import (
	"net/http"

	"github.com/clarke94/weigh-me/srv/internal/driver"
	"github.com/go-chi/render"
)

// GetAllWeight will return all weight values in database
func GetAllWeight(w http.ResponseWriter, r *http.Request) {
	weights, err := driver.GetAllWeights()

	if err != nil {
		http.Error(w, "Error While Getting All Weights", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, weights)
}
