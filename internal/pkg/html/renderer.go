package html

import (
	"fmt"
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

func (r *Renderer) RenderTemplate(w http.ResponseWriter, template string, args ...any) {
	data := make(map[string]any)

	if len(args)%2 != 0 {
		http.Error(w,
			"args must be even number of elements: key, value...",
			http.StatusInternalServerError,
		)
		return
	}

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			http.Error(w,
				fmt.Sprintf("argument %d is not a string key", i),
				http.StatusInternalServerError,
			)
			return
		}

		data[key] = args[i+1]
	}

	if err := r.views.ExecuteTemplate(w, template, data); err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}

	w.WriteHeader(http.StatusOK)
}
