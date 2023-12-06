package web

import (
	"log"
	"net/http"
)

type Handler struct {
	Logger  *log.Logger
	handler http.Handler
}

func (h *Handler) init() {
	//we will initialize all the routes here
	r := mux.NewRouter()
	h.handler = r
}
