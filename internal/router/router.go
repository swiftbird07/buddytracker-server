package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swiftbird07/buddytracker-server/internal/handler"
	mdlwr "github.com/swiftbird07/buddytracker-server/internal/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", handler.Register)                          // Register
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {})   // Login
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {}) // Terminate session
		})

		r.Route("/", func(r chi.Router) {
			r.Use(mdlwr.Auth)

			r.Route("/buddies", func(r chi.Router) {
				r.Get("/", handler.ListBuddies)

				r.Route("/{user-id}", func(r chi.Router) {
					r.Delete("/", func(w http.ResponseWriter, r *http.Request) {}) // Friends
				})
			})

			r.Route("/users", func(r chi.Router) {
				r.Post("/", func(w http.ResponseWriter, r *http.Request) {}) // (By code)
				r.Route("/{user-id}", func(r chi.Router) {
					r.Get("/", func(w http.ResponseWriter, r *http.Request) {})      // Friends
					r.Put("/", func(w http.ResponseWriter, r *http.Request) {})      // Self
					r.Delete("/", func(w http.ResponseWriter, r *http.Request) {})   // Self
					r.Get("/code", func(w http.ResponseWriter, r *http.Request) {})  // Self
					r.Post("/code", func(w http.ResponseWriter, r *http.Request) {}) // Self

					r.Route("/status", func(r chi.Router) {
						r.Get("/", func(w http.ResponseWriter, r *http.Request) {}) // Friends
						r.Put("/", func(w http.ResponseWriter, r *http.Request) {}) // Self
					})
				})
			})

			r.Route("/activities", func(r chi.Router) {
				r.Get("/", handler.ListActivities)
			})
		})
	})

	return r
}
