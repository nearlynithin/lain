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

	r.Handle("/user-follows", mux.MethodHandler{
		http.MethodPost:   h.followUser,
		http.MethodDelete: h.unfollowUser,
	})

	r.Handle("/*", mux.MethodHandler{
		http.MethodGet: h.static(),
	})

	gob.Register(lain.User{})
	gob.Register(url.Values{})
	h.session = sessions.New(h.SessionKey)

	//this is a list of middlewares
	h.handler = r
	h.handler = h.withUser(h.handler)
	h.handler = h.session.Enable(h.handler)
	h.handler = withMethodOverride(h.handler)
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

// code from Alex edwards
// This is "Middleware", that checks if the request method is a post request
func withMethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only act on POST requests.
		if r.Method == "POST" {

			// Look in the request body and headers for a spoofed method.
			// Prefer the value in the request body if they conflict.
			method := r.PostFormValue("_method")

			// Check that the spoofed method is a valid HTTP method and
			// update the request object accordingly.
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
