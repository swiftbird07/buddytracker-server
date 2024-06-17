package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func Logger(next http.Handler) http.Handler {
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

		_, isDebug := os.LookupEnv("DEBUG")
		if !isDebug {
			// Log remote Ip and real IP of the client if they differ so analysis can be done if needed
			if sourceIP != r.RemoteAddr {
				log.Printf("Buddy Tracker Server - Parsed IP Information - Remote Connection %s has real IP: %s\n", r.RemoteAddr, sourceIP)
			}

			next.ServeHTTP(w, r)
			return
		}

		bHeader, err := json.MarshalIndent(r.Header, "", "  ")
		if err != nil {
			return
		}

		bCookies, err := json.MarshalIndent(r.Cookies(), "", "  ")
		if err != nil {
			return
		}
		log.SetOutput(os.Stdout)

		log.Printf("\nRemoteAddr: %s\nParsed IP: %s\nURL: %s\nHeader: %s\nCookies: %s\n", r.RemoteAddr, sourceIP, r.URL.String(), string(bHeader), string(bCookies))

		next.ServeHTTP(w, r)
	})
}
