package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/swiftbird07/buddytracker-server/internal/controller"
)

type Key string

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get source IP for logging
		var sourceIP string
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			sourceIP = strings.Split(forwardedFor, ",")[0]
		} else if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
			sourceIP = realIP
		} else {
			sourceIP = r.RemoteAddr
		}

		// Check if User Agent starts with "Buddy Tracker"
		if !strings.HasPrefix(r.UserAgent(), "Buddy Tracker") {
			log.Println("Buddy Tracker Server - Auth Fail - User-Agent mismatch for incoming connection (Source IP: ", sourceIP+". User-Agent: ", r.UserAgent()+").")
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// Check if Authorization header is present
		authHeader := r.Header.Get("Authorization")
		token, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found {
			log.Println("Buddy Tracker Server - Auth Fail - No token found in Authorization header for incoming connection (Source IP: ", sourceIP+". User-Agent: ", r.UserAgent()+").")
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// Validate token
		validSession, user := controller.ValidToken(token)

		if !validSession {
			log.Println("Buddy Tracker Server - Auth Fail - Authentication failed for incoming connection (Source IP: ", sourceIP+". User-Agent: ", r.UserAgent()+").")
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		} else {
			// Get username
			log.Println("Buddy Tracker Server - Auth Success - Authentication successful for incoming connection. (User: ", user.Name+", Source IP: ", sourceIP+". User-Agent: ", r.UserAgent()+").")
			ctx := context.WithValue(r.Context(), Key("user"), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
