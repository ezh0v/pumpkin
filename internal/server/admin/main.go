package admin

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/server/pkg/response"
)

func home(app *app.Context) http.HandlerFunc {
	page := renderer.NewPage("home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		response.WithPage(w, page)
	}
}
