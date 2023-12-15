package web

import (
	"errors"
	"net/http"
	"net/url"

	"lain.sceptix.net"
)

var loginTmpl = parseTmpl("login.tmpl")

type loginData struct {
	Form url.Values
	Err  error
}

// this is a method inside of that struct, it is used to render the login page
func (h *Handler) renderLogin(w http.ResponseWriter, data loginData, statusCode int) {
	//we'll have the renderTmpl function ready at this point
	h.renderTmpl(w, loginTmpl, data, statusCode)
}

func (h *Handler) showLogin(w http.ResponseWriter, r *http.Request) {
	//This is gonna call the renderLogin function
	h.renderLogin(w, loginData{}, http.StatusOK)
	//and we need to register this handker into the router
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	//first we pass the form we et from login page tmpl
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	input := lain.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formPtr(r.PostForm, "username"),
	}

	user, err := h.Service.Login(ctx, input)
	if errors.Is(err, lain.ErrUserNotFound) || errors.Is(err, lain.ErrUsernameTaken) {
		h.renderLogin(w, loginData{
			Form: r.PostForm,
			Err:  err,
		}, http.StatusBadRequest)
		return
	}
	if err != nil {
		h.Logger.Printf("could not login: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	//putting the user in the session using the user kkey
	h.session.Put(r, "user", user)
	//redirecting back to the homepage
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	h.session.Remove(r, "user")
	http.Redirect(w, r, "/", http.StatusFound)
}

func formPtr(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}

	s := form.Get(key)
	return &s
}
