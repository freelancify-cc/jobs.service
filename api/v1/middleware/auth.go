package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/freelancify/jobs/config"
	"github.com/freelancify/jobs/helpers"
)

func EnsureAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		token := ""
		cookieToken, err := r.Cookie("token")
		if err != nil {
			if len(authHeader) != 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			token = authHeader[1]
		} else {
			token = cookieToken.String()
		}

		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// extract which user is making the request by calling auth userinfo endpoint
func ExtractUserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("token").(string)
		if token == "" {
			http.Error(w, "token is empty", http.StatusInternalServerError)
			return
		}

		client := http.Client{}
		req, err := http.NewRequest("GET", config.GetConfig().AuthServiceUrl+"/api/auth/userinfo", nil)
		if err != nil {
			http.Error(w, "Could not reach auth service", http.StatusInternalServerError)
			return
		}
		req.Cookie(token)
		req.Header = http.Header{
			"Authorization": {"Bearer " + token},
		}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if res.StatusCode != 200 {
			http.Error(w, "could not fetch userinfo", http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()
		jsonRes, err := helpers.ParseJsonBody(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userId := jsonRes["sub"].(uuid.UUID)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
