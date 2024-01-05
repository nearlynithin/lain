package web

import (
	"errors"
	"net/http"
	"net/url"

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

func (h *Handler) putErr(r *http.Request, key string, err error) {
	h.session.Put(r, key, err.Error())
}

func (h *Handler) popErr(r *http.Request, key string) error {
	s := h.session.PopString(r, key)

	if s != "" {
		return errors.New(s)
	}
	return nil
}

func (h *Handler) popForm(r *http.Request, key string) url.Values {
	v, _ := h.session.Pop(r, key).(url.Values)
	return v
}
