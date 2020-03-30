package user

import (
	"github.com/go-chi/chi"
)

// Routes holds all endpoints for the user REST API
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetUserByAuthID)
	router.Post("/", InsertUser)
	return router
}
