package main

import (
	"log"
	"net/http"
	"time"

    "github.com/freelancing/jobs/api/v1/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func initializeBookclubRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("bookclub server"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", router.SetupRoutes())
	})
}

func initializeBookclubServer() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)

	r.Use(middleware.Timeout(20 * time.Second))

	initializeBookclubRoutes(r)

	return r
}

func main() {
	// load config file
	//config := config.GetConfig()

	// initialize database
	//database.Initialize(config)


	// initialize server
	r := initializeBookclubServer()

    log.Printf("[!] Starting freelancify job service on %s:%s", "127.0.0.1", "4444")

	// start listener
    http.ListenAndServe("127.0.0.1"+ ":" + "4444", r)
}
