package admin

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/server/handlers"
	"github.com/ezh0v/pumpkin/internal/server/response"
)

func home(c *handlers.Context) http.HandlerFunc {
	page := c.NewPage("home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		response.WithPage(w, page)
	}
}
