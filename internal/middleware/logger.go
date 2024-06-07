package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, isDebug := os.LookupEnv("DEBUG")
		if !isDebug {
			next.ServeHTTP(w, r)
			return
		}

		bHeader, err := json.MarshalIndent(r.Header, "", "  ")
		if err != nil {
			return
		}

		bData, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		bCookies, err := json.MarshalIndent(r.Cookies(), "", "  ")
		if err != nil {
			return
		}
		log.SetOutput(os.Stdout)

		log.Printf("\nRemoteAddr: %s\nURL: %s\nHeader: %s\nCookies: %s\nBody: %s\n", r.RemoteAddr, r.URL.String(), string(bHeader), string(bCookies), string(bData))

		next.ServeHTTP(w, r)
	})
}
