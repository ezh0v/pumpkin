package web

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/renderer"
)

func home(app *app.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := renderer.RenderTemplate(w, templates, "home.html", "app", app); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
