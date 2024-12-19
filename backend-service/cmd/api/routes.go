package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	mux.Use(csrfMiddleware)
	mux.Use(app.CRSFAdder)

	
	mux.Use(app.authenticate)
	
	mux.Post("/create-user", app.CreateUserHandler)
	mux.Post("/login", app.Login)
	mux.Get("/refresh", app.RefreshToken)
	mux.Get("/logout", app.Logout)
	
	mux.Post("/create-ai-article", app.CreateAiArticle)

	mux.Get("/news/{page}", app.GetNews)
	mux.Get("/article/{id}", app.GetArticle)

	mux.Route("/user", func(r chi.Router) {
		r.Use(app.requireAuthenticatedUser)

		r.Get("/data", app.GetUserData)
		r.Get("/check-admin", app.GetCheckAdmin)
	})

	mux.Route("/admin", func(r chi.Router) {
		r.Post("/create-article", app.CreateNewArticle)
	})

	return mux
}
