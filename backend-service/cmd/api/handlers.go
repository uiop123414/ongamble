package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"ongambl/internal/models"
	"ongambl/internal/validator"
	"slices"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	v := validator.New()

	models.ValidateEmail(v, input.Email)
	models.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		app.errorJSONWithMSG(w, errors.New("invalid credentials"), v.Errors, http.StatusUnprocessableEntity)
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	err = user.SetPasswordHash()

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.DB.CreateUser(user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Successfully Created",
	}

	app.writeJSON(w, http.StatusCreated, resp)
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	v := validator.New()

	if models.ValidatePasswordPlaintext(v, input.Password); !v.Valid() {
		app.errorJSONWithMSG(w, errors.New("invalid credentials"), v.Errors, http.StatusUnprocessableEntity)
		return
	}

	user, err := app.DB.GetUserByUsername(input.Username)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	isMatches, err := user.PasswordMatches(input.Password)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Successfully login",
	}

	if !isMatches {
		resp.Error = true
		resp.Message = "Invalid password or username"
		app.writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	u := jwtUser{
		ID:       user.ID,
		Username: user.Username,
	}

	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)

	http.SetCookie(w, refreshCookie)

	resp.Data = tokens

	app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) GetUserData(w http.ResponseWriter, r *http.Request) {

	token, err := app.GetAuthToken(r)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := app.DB.GetUserByToken(models.ScopeAuthentication, token)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := JSONResponse{
		Error:   false,
		Message: "User Data",
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) RefreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			_, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(app.cfg.jwt.secret), nil
			})
			if err != nil {
				http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserByID(int64(userID))
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:       user.ID,
				Username: user.Username,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			err = app.DB.DeleteAllTokensForUser(models.ScopeAuthentication, int64(user.ID))
			if err != nil {
				app.errorJSON(w, errors.New("server error"), http.StatusUnauthorized)
				return
			}

			tk := models.Token{
				Hash:   []byte(tokenPairs.Token),
				UserID: user.ID,
				Expiry: time.Now().Add(app.auth.TokenExpiry),
				Scope:  models.ScopeAuthentication,
			}

			err = app.DB.InsertToken(&tk)

			if err != nil {
				fmt.Println(err)
				app.errorJSON(w, errors.New("server error"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {

}
func (app *application) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ArticleName string        `json:"articleName"`
		Username    string        `json:"username"`
		Time        string        `json:"time"`
		Blocks      []interface{} `json:"blocks"`
		Publish     bool          `json:"publish"`
		Version     string        `json:"version"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	htmlList, err := json.Marshal(input.Blocks)
	if err != nil {
		fmt.Println("Error marshaling Blocks:", err)
		return
	}

	blocksString := string(htmlList)

	article := models.Article{
		Name:        input.ArticleName,
		Username:    input.Username,
		ReadingTime: input.Time,
		HtmlList:    blocksString,
		Publish:     input.Publish,
	}

	err = app.DB.NewArticle(&article)
	fmt.Println(err)
	payload := JSONResponse{
		Error:   false,
		Message: "Article susscefully created",
	}

	if err != nil {
		payload.Error = true
		payload.Message = err.Error()
		app.writeJSON(w, http.StatusBadRequest, payload)
	}

	app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) GetArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	article, err := app.DB.GetArticle(id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	payload := JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("Article №%d", id),
		Data:    article,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) GetNews(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	news, err := app.DB.GetNews(page)
	if err != nil {
		fmt.Print(news)
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	payload := JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("News in page: %d", page),
		Data:    news,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) GetCheckAdmin(w http.ResponseWriter, r *http.Request) {
	token, err := app.GetAuthToken(r)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}

	permissions, err := app.DB.GetUserPermissions(token)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := JSONResponse{
		Error:   false,
		Message: "User is admin",
		Data:    true,
	}

	if slices.Contains(*permissions, models.AdminWrite){
		app.writeJSON(w, http.StatusOK, payload)
	} else {
		payload.Message = "User is not admin"
		payload.Data = false
		app.writeJSON(w, http.StatusOK, payload)
	}
}