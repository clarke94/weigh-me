package weight

import (
	"net/http"

	"github.com/go-chi/render"
)

// GetAllWeight will return all weight values in database
func GetAllWeight(w http.ResponseWriter, r *http.Request) {
	weights := []Weight{
		{
			ID:    "001",
			Name:  "Wed Feb 26 2020 00:00:00 GMT+0000 (Greenwich Mean Time)",
			Value: 80,
		},
	}

	render.JSON(w, r, weights)
}
