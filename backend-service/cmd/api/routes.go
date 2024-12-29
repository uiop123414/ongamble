package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Global middleware applied to all routes
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// Grafana metrics route
	mux.Route("/metrics", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			promhttp.Handler().ServeHTTP(w, r)
		})
	})

	// Main application routes under "/"
	appRouter := chi.NewRouter()

	// Middleware for main app routes
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	appRouter.Use(csrfMiddleware)
	appRouter.Use(app.CRSFAdder)
	appRouter.Use(app.authenticate)

	// Authenticated routes
	appRouter.Post("/create-user", app.CreateUserHandler)
	appRouter.Post("/login", app.Login)
	appRouter.Get("/refresh", app.RefreshToken)
	appRouter.Get("/logout", app.Logout)
	appRouter.Post("/create-ai-article", app.CreateAiArticle)
	appRouter.Get("/news/{page}", app.GetNews)
	appRouter.Get("/article/{id}", app.GetArticle)

	// Protected user routes
	appRouter.Route("/user", func(r chi.Router) {
		r.Use(app.requireAuthenticatedUser)
		r.Get("/data", app.GetUserData)
		r.Get("/check-admin", app.GetCheckAdmin)
	})

	// Protected admin routes
	appRouter.Route("/admin", func(r chi.Router) {
		r.Use(app.requireAuthenticatedUser) // Only allow authenticated users
		r.Post("/create-article", app.CreateNewArticle)
	})

	// Mount appRouter under "/"
	mux.Mount("/", appRouter)

	return mux
}
