package weight

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetByID will get one weight value by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")

	weights := Weight{
		ID:    ID,
		Name:  "Wed Feb 26 2020 00:00:00 GMT+0000 (Greenwich Mean Time)",
		Value: 80,
	}

	render.JSON(w, r, weights)
}
