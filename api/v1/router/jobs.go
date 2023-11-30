package router

import (
	jobshandler "github.com/freelancify/jobs/api/v1/handlers"
	"github.com/freelancify/jobs/api/v1/middleware"
	"github.com/go-chi/chi"
)

type JobRoutes struct{}

func (j JobRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{", func(jobRoutes chi.Router) {
		jobRoutes.Use(middleware.EnsureAuth)
		jobRoutes.Use(middleware.ExtractUserId)
		jobRoutes.Get("/", jobshandler.GetJobDetails)
		jobRoutes.Get("/", jobshandler.GetAllJobs)
		jobRoutes.Post("/", jobshandler.CreateJob)
	})

	return r
}
