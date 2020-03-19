package weight

import (
	"net/http"

	"github.com/go-chi/render"
)

// CreateWeight will create one weight data for a given user
func CreateWeight(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Weight successfully"

	render.JSON(w, r, response)
}
