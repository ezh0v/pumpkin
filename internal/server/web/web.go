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
		app.RenderHTML(http.StatusOK, w, "home.html", map[string]any{
			"appVersion": app.Version,
		})
	}
}
