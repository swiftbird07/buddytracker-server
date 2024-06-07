package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/swiftbird07/buddytracker-server/internal/controller"
)

type Key string

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		validSession, user := controller.ValidToken(token)

		if !validSession {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		} else {
			ctx := context.WithValue(r.Context(), Key("user"), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
