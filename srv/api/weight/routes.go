package weight

import (
	"github.com/go-chi/chi"
)

// Routes holds all endpoints for weight REST API
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{ID}", GetByID)
	router.Delete("/{ID}", DeleteByID)
	router.Post("/", CreateWeight)
	router.Get("/", GetAllWeight)
	return router
}
