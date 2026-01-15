package web

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/renderer"
)

func home(ctx *app.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := renderer.RenderTemplate(w, templates, "home.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
