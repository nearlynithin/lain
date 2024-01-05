package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"mvdan.cc/xurls/v2"
)

//go:embed template/include/*.tmpl template/*.tmpl
var templateFS embed.FS

var tmplFuncs = template.FuncMap{
	"linkify": linkify,
}

var reURL = xurls.Relaxed()

func linkify(s string) template.HTML {
	s = template.HTMLEscapeString(s)
	return template.HTML(reURL.ReplaceAllString(s, `<a href="$0" target="_blank" rel="noopener noreferrer">$0</a>`))
}

func parseTmpl(name string) *template.Template {
	tmpl := template.New(name).Funcs(tmplFuncs)
	tmpl = template.Must(tmpl.ParseFS(templateFS, "template/include/*.tmpl"))
	return template.Must(tmpl.ParseFS(templateFS, "template/"+name))
}

// A utility function for renderlogin because we'll be rendering a lot of templates
func (h *Handler) renderTmpl(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	//we'll create a "Buffer" here to render the template first and then write the response back
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, data)
	if err != nil {
		//We use a logger here too cuz this might return something that user is not supposed to see ig
		h.Logger.Output(2, fmt.Sprintf("could not render %q: %v\n", tmpl.Name(), err))
		http.Error(w, fmt.Sprintf("could not render %q", tmpl.Name()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = buff.WriteTo(w)
	if err != nil {
		h.Logger.Output(2, fmt.Sprintf("could not send %q: %v\n", tmpl.Name(), err))
	}

}
