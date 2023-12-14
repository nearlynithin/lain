package web

import (
	"net/http"

	"lain.sceptix.net"
)

type Session struct {
	IsLoggedIn bool
	User       lain.User
}

func (h *Handler) sessionFromReq(r *http.Request) Session {
	var out Session

	if h.session.Exists(r, "user") {
		user, ok := h.session.Get(r, "user").(lain.User)
		out.IsLoggedIn = ok
		out.User = user
	}

	return out
}
