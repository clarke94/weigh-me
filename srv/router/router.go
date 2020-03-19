package router

import (
	"github.com/clarke94/weigh-me/srv/api/weight"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes holds all API routes to server
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/weight", weight.Routes())
	})

	return router
}
