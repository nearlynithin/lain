package web

import (
	"net/http"
	"net/url"

	"lain.sceptix.net"
)

var homeTmpl = parseTmpl("home.tmpl")

type homeData struct {
	Session
	CreatePostErr  error
	CreatePostForm url.Values
	Posts          []lain.PostsRow
}

func (h *Handler) renderHome(w http.ResponseWriter, data homeData, statusCode int) {
	h.renderTmpl(w, homeTmpl, data, statusCode)
}

func (h *Handler) showHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pp, err := h.Service.Posts(ctx)
	if err != nil {
		h.log(err)
		h.renderErr(w, err)
		return
	}

	h.renderHome(w, homeData{
		Session:        h.sessionFromReq(r),
		CreatePostErr:  h.popErr(r, "create_post_err"),
		CreatePostForm: h.popForm(r, "create_post_form"),
		Posts:          pp,
	}, http.StatusOK)
}
