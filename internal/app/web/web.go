package web

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

func Route(c *app.Context) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", home(c))
	return handler
}

func home(c *app.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.RenderHTML(http.StatusOK, w, "home.html", nil)
	}
}
