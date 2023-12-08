package web

import (
	"log"
	"net/http"
	"sync"
	"github.com/nicolasparada/go-mux"
)

type Handler struct {
	Logger  *log.Logger
	once    sync.Once
	handler http.Handler
}

func (h *Handler) init() {
	//we will initialize all the routes here
	r := mux.NewRouter()
	r.Handle("/login", mux.MethodHandler{
		http.MethodGet: h.showLogin,
	})

	h.handler = r
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//to initialize all the routes
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
