package main

import (
	"errors"
	"net/http"
	"ongambl/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/csrf"
	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})

}

func (app *application) CRSFAdder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			r = app.contextSetUser(r, models.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.cfg.jwt.secret))
		if err != nil {
			app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			return
		}

		if claims.Issuer != app.cfg.jwt.issuer {
			app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			return
		}

		if !claims.AcceptAudience(app.cfg.jwt.audience) {
			app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		user, err := app.DB.GetUserByID(userID)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrRecordNotFound):
				app.errorJSON(w, errors.New("invalid authentication token"), http.StatusUnauthorized)
			default:
				app.errorJSON(w, err)

			}
			return
		}

		r = app.contextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		if user.IsAnonymous() {
			app.errorJSON(w, errors.New("authantication required"), http.StatusProxyAuthRequired)
			return
		}

		next.ServeHTTP(w, r)
	})
}
