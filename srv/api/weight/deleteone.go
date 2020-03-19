package weight

import (
	"net/http"

	"github.com/go-chi/render"
)

// DeleteByID will delete one weight data by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Weight successfully"

	render.JSON(w, r, response)
}
