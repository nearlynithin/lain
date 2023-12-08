package web

import (
	"net/http"
	"net/url"
)

var loginTmpl = parseTmpl("login.tmpl")

type loginData struct {
	Form url.Values
	Err  error
}

// this is a method inside of that struct, it is used to render the login page
func (h *Handler) renderLogin(w http.ResponseWriter, data loginData, statusCode int){
      //we'll have the renderTmpl function ready at this point
	h.renderTmpl(w,loginTmpl,data,statusCode)
}

func(h *Handler) showLogin( w http.ResponseWriter, r *http.Request){
	//This is gonna call the renderLogin function	 
	h.renderLogin(w,loginData{},http.StatusOK)
}