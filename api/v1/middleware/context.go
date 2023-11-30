package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func JobCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("job ctx middleware called")
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
