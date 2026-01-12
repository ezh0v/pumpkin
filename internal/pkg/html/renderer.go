package html

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	views *template.Template
}

func NewRenderer(views *template.Template) Renderer {
	return Renderer{
		views: views,
	}
}

func (r *Renderer) RenderHTML(statusCode int, w http.ResponseWriter, name string, data any) {
	w.WriteHeader(statusCode)
	if err := r.views.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
