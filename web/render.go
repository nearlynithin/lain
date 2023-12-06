package web

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

var templateFS embed.FS

// here we created a function to parse the templates like login and layout stuff
func parseTmpl(name string) *template.Template {
	tmpl := template.New(name)
	tmpl = template.Must(template.ParseFS(templateFS, "template/include/*.html"))
	return template.Must(tmpl.ParseFS(templateFS, "template/"+name))
}

// A utility function for renderlogin because we'll be rendering a lot of templates
func (h *Handler) renderTmpl(w http.ResponseWriter, tmpl *template.Template, data any, statusCode int) {
	//we'll create a "Buffer" here to render the template first and then write the response back
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, data)
	if err != nil {
		//We use a logger here too cuz this might return something that user is not supposed to see ig
		h.Logger.Output(2, fmt.Sprintf("could not render %q: %v\n", tmpl.Name, err))
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
