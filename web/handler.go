package web

import (
	"encoding/gob"
	"log"
	"net/http"
	"sync"

	"github.com/golangcollege/sessions"
	"github.com/nicolasparada/go-mux"
	"lain.sceptix.net"
)

type Handler struct {
	Logger     *log.Logger
	Service    *lain.Service
	SessionKey []byte
	once       sync.Once
	handler    http.Handler
	session    *sessions.Session
}

func (h *Handler) init() {
	//we will initialize all the routes here
	r := mux.NewRouter()

	r.Handle("/", mux.MethodHandler{
		http.MethodGet: h.showHome,
	})

	r.Handle("/login", mux.MethodHandler{
		http.MethodGet:  h.showLogin,
		http.MethodPost: h.login,
	})

	r.Handle("/logout", mux.MethodHandler{
		http.MethodPost: h.logout,
	})

	gob.Register(lain.User{})

	h.session = sessions.New(h.SessionKey)
	h.handler = r
	h.handler = h.session.Enable(h.handler)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//to initialize all the routes
	//calling init only once
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
