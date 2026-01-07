package html

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	Views *template.Template
}

func (r *Renderer) RenderHTML(statusCode int, w http.ResponseWriter, name string, data any) {
	w.WriteHeader(statusCode)
	if err := r.Views.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
