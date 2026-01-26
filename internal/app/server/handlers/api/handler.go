package api

import (
	"net/http"

	"github.com/ezh0v/pumpkin/internal/app"
)

func Handler(app *app.Instance) http.Handler {
	handler := http.NewServeMux()
	return http.StripPrefix("/api", handler)
}
