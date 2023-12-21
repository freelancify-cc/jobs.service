package router

import (
	jobshandler "github.com/freelancify/jobs/api/v1/handlers"
	"github.com/freelancify/jobs/api/v1/middleware"
	"github.com/go-chi/chi"
)

type JobRoutes struct{}

func (j JobRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", jobshandler.GetAllJobs)
	r.Get("/{id}", jobshandler.GetJobDetails)

	r.Group(func(routes chi.Router) {
		routes.Use(middleware.EnsureAuth)
		routes.Use(middleware.ExtractUserId)
		routes.Post("/", jobshandler.CreateJob)
	})

	/*
		r.Group(func(routes chi.Router) {
			routes.Use(middleware.EnsureAuth)
			routes.Use(middleware.ExtractUserId)
			routes.Use(middleware.JobCtx)
		})


			r.Route("/", func(jobRoutes chi.Router) {
				jobRoutes.Get("/", jobshandler.GetAllJobs)
				jobRoutes.Route("/create", func(postJobRoute chi.Router) {
					postJobRoute.Use(middleware.EnsureAuth)
					postJobRoute.Use(middleware.ExtractUserId)

				})
			})
	*/
	return r
}
