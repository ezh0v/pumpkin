package web

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

func Route(app *app.Context) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", home(app))
	return handler
}

func home(app *app.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.RenderTemplate(w, "home.html",
			"appVersion", app.Version,
		)
	}
}
