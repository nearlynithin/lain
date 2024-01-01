package web

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/nicolasparada/go-errs/httperrs"
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
		h.renderLogin(w, loginData{Err: errors.New("bad request")}, http.StatusBadRequest)
	}

	ctx := r.Context()
	input := lain.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formPtr(r.PostForm, "username"),
	}

	user, err := h.Service.Login(ctx, input)
	if err != nil {
		h.log(err)
		h.renderLogin(w, loginData{
			Form: r.PostForm,
			Err:  maskErr(err),
		}, httperrs.Code(err))
		return
	}

	//putting the user in the session using the user kkey
	h.session.Put(r, "user", user)
	//redirecting back to the homepage
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) log(err error) {
	if httperrs.IsInternalServerError(err) {
		_ = h.Logger.Output(2, err.Error())
	}
}

func maskErr(err error) error {
	if httperrs.IsInternalServerError(err) {
		return errors.New("internal server error")
	}
	return err
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

// Middleware
func (h *Handler) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !h.session.Exists(r, "user") {
			next.ServeHTTP(w, r)
			return
		}

		usr, ok := h.session.Get(r, "user").(lain.User)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		ctx = lain.ContextWithUser(ctx, usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
