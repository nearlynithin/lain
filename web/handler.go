package web

import (
	"embed"
	"encoding/gob"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/golangcollege/sessions"
	"github.com/nicolasparada/go-mux"
	"lain.sceptix.net"
)

//go:embed all:static
var staticFS embed.FS

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

	r.Handle("/posts", mux.MethodHandler{
		http.MethodPost: h.createPost,
	})

	r.Handle("/p/{postID}", mux.MethodHandler{
		http.MethodGet: h.showPost,
	})

	r.Handle("/comments", mux.MethodHandler{
		http.MethodPost: h.createComment,
	})

	//cool
	r.Handle("/@{username}", mux.MethodHandler{
		http.MethodGet: h.showUser,
	})

	r.Handle("/*", mux.MethodHandler{
		http.MethodGet: h.static(),
	})

	gob.Register(lain.User{})
	gob.Register(url.Values{})
	h.session = sessions.New(h.SessionKey)

	h.handler = r
	h.handler = h.withUser(h.handler)
	h.handler = h.session.Enable(h.handler)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//to initialize all the routes
	//calling init only once
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}

func (h *Handler) static() http.HandlerFunc {
	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(sub)).ServeHTTP
}
