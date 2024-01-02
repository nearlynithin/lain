package web

import (
	"net/http"

	"github.com/nicolasparada/go-errs/httperrs"
)

var errTmpl = parseTmpl("err.tmpl")

type errData struct {
	Err error
}

func (h *Handler) renderErr(w http.ResponseWriter, err error) {
	h.renderTmpl(w, errTmpl, errData{
		Err: maskErr(err),
	}, httperrs.Code(err))
}
