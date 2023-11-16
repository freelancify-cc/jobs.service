package router

import (
	"github.com/go-chi/chi"
)

type JobRoutes struct{}

func (j JobRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(resourceRequestRoutes chi.Router) {
	})

	r.Route("/", func(rootRequestRoutes chi.Router) {
	})

	return r
}
