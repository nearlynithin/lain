package web

import (
	"net/http"

	"github.com/nicolasparada/go-mux"
	"lain.sceptix.net"
)

// parsing the user template
var userTmpl = parseTmpl("user.tmpl")

type userData struct {
	Session
	User  lain.User
	Posts []lain.PostsRow
}

// a function to render this template
func (h *Handler) renderUser(w http.ResponseWriter, data userData, statusCode int) {
	h.renderTmpl(w, userTmpl, data, statusCode)
}

func (h *Handler) showUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := mux.URLParam(ctx, "username")

	//we'll have the username here, so calling the service
	usr, err := h.Service.UserByUsername(ctx, username)
	if err != nil {
		h.log(err)
		h.renderErr(w, r, err)
		return
	}

	//returning the posts here BUT with a username, so only that users' posts are returned like defined in the posts function
	pp, err := h.Service.Posts(ctx, username)
	if err != nil {
		h.log(err)
		h.renderErr(w, r, err)
		return
	}

	//calling the posts renderer defined above
	h.renderUser(w, userData{
		Session: h.sessionFromReq(r),
		User:    usr,
		Posts:   pp,
	}, http.StatusOK)

}
