package main

import (
	"log"
	"net/http"
	"time"

    "github.com/freelancing/jobs/config"
    "github.com/freelancing/jobs/api/v1/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initializeRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("bookclub server"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", router.SetupRoutes())
	})
}

func initializeServer() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)

	r.Use(middleware.Timeout(20 * time.Second))

	initializeRoutes(r)

	return r
}

func main() {
	// load config file
	config := config.GetConfig()

	// initialize database
	//database.Initialize(config)


	// initialize server
	r := initializeServer()

    log.Printf("[!] Starting freelancify job service on %s:%s", config.Host, config.Port)

	// start listener
    http.ListenAndServe(config.Host+ ":" + config.Port, r)
}
