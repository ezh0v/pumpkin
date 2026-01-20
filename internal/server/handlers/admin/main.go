package admin

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/server/handlers"
	"github.com/ezh0v/pumpkin/internal/server/pkg/response"
)

func home(context *handlers.Context) http.HandlerFunc {
	page := context.Renderer.NewPage("home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		response.WithPage(w, page)
	}
}
