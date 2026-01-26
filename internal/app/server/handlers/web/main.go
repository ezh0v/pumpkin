package web

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app/server/handlers"
	"github.com/ezh0v/pumpkin/internal/app/server/response"
)

func home(c *handlers.Context) http.HandlerFunc {
	page := c.NewPage("home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		response.WithPage(w, page)
	}
}
