package api

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

func Route(app *app.Context) http.Handler {
	handler := http.NewServeMux()
	return http.StripPrefix("/api", handler)
}
